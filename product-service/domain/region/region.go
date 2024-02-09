package region

import (
	"context"
	"fmt"
	"reflect"

	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Region, error)
	Get(ctx context.Context, filter entity.Region) (entity.Region, error)
	Create(ctx context.Context, region entity.Region) (entity.Region, error)
	Update(ctx context.Context, region entity.Region) (entity.Region, error)
	Delete(ctx context.Context, region entity.Region) error
}

type region struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
	cache      cache.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection, cache cache.Interface) Interface {
	return &region{
		logger:     logger,
		json:       json,
		collection: db,
		cache:      cache,
	}
}

// List returns list of regions
func (r *region) List(ctx context.Context) ([]entity.Region, error) {
	regions := []entity.Region{}
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return regions, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &regions); err != nil {
		return regions, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return regions, nil
}

// Get returns specific region
func (r *region) Get(ctx context.Context, req entity.Region) (entity.Region, error) {
	var region entity.Region
	var filter any
	var cacheKey string

	r.logger.Debug(ctx, req)
	switch {
	case !req.Id.IsZero():
		filter = bson.M{"_id": req.Id}
		cacheKey = fmt.Sprintf("id:%v", req.Id)
	case req.Name != "":
		filter = bson.M{"name": req.Name}
		cacheKey = fmt.Sprintf("name:%v", req.Name)
	}

	// get from cache, if no error and region found, direct return
	region, err := r.getCache(ctx, fmt.Sprintf(entity.CacheKeyRegion, cacheKey))
	if err == nil && !region.Id.IsZero() {
		r.logger.Info(ctx, fmt.Sprintf("cache for %v found", cacheKey))
		return region, nil
	} else if err != nil {
		r.logger.Error(ctx, err)
	}
	r.logger.Info(ctx, fmt.Sprintf("cache for %v no found", cacheKey))

	if err := r.collection.FindOne(ctx, filter).Decode(&region); err != nil && err == mongo.ErrNoDocuments {
		return region, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return region, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	// set region cache if result found from mongo
	if err := r.setCache(ctx, fmt.Sprintf(entity.CacheKeyRegion, cacheKey), region); err != nil {
		r.logger.Error(ctx, fmt.Sprintf("cache for region:%v failed to be set", req.Id.Hex()))
	}

	return region, nil
}

// Create creates new data
func (r *region) Create(ctx context.Context, region entity.Region) (entity.Region, error) {
	res, err := r.collection.InsertOne(ctx, region)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return region, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return region, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newregion, err := r.Get(ctx, entity.Region{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newregion, err
	}

	return newregion, nil
}

// Update updates existing data
func (r *region) Update(ctx context.Context, region entity.Region) (entity.Region, error) {
	filter := bson.M{"_id": region.Id}

	// Prepare the update document with non-zero values only
	update := bson.M{"$set": bson.M{}}

	// Helper function to add a field to the update document if it has a non-zero value
	addField := func(fieldName string, fieldValue interface{}) {
		if !reflect.ValueOf(fieldValue).IsZero() {
			update["$set"].(bson.M)[fieldName] = fieldValue
		}
	}

	addField("name", region.Name)

	// If no non-zero fields, return the original region without updating
	if len(update["$set"].(bson.M)) == 0 {
		return r.Get(ctx, entity.Region{Id: region.Id})
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return region, errors.NewWithCode(codes.CodeNoSQLUpdate, err.Error())
	}

	newregion, err := r.Get(ctx, entity.Region{Id: region.Id})
	if err != nil {
		return newregion, err
	}

	return newregion, nil
}

// Delete deletes existing data
func (r *region) Delete(ctx context.Context, region entity.Region) error {
	filter := bson.M{"_id": region.Id}

	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewWithCode(codes.CodeNoSQLDelete, err.Error())
	}

	if res.DeletedCount < 1 {
		return errors.NewWithCode(codes.CodeNoSQLNoRowsAffected, fmt.Sprintf("failed to delete id %v", region.Id.Hex()))
	}

	return nil
}
