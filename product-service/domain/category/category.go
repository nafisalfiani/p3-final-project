package category

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
	List(ctx context.Context) ([]entity.Category, error)
	Get(ctx context.Context, filter entity.Category) (entity.Category, error)
	Create(ctx context.Context, category entity.Category) (entity.Category, error)
	Update(ctx context.Context, category entity.Category) (entity.Category, error)
	Delete(ctx context.Context, category entity.Category) error
}

type category struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
	cache      cache.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection, cache cache.Interface) Interface {
	return &category{
		logger:     logger,
		json:       json,
		collection: db,
		cache:      cache,
	}
}

// List returns list of categories
func (c *category) List(ctx context.Context) ([]entity.Category, error) {
	categories := []entity.Category{}
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return categories, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &categories); err != nil {
		return categories, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return categories, nil
}

// Get returns specific category
func (c *category) Get(ctx context.Context, req entity.Category) (entity.Category, error) {
	var category entity.Category
	var filter any
	var cacheKey string

	c.logger.Debug(ctx, req)
	switch {
	case !req.Id.IsZero():
		filter = bson.M{"_id": req.Id}
		cacheKey = fmt.Sprintf("id:%v", req.Id)
	case req.Name != "":
		filter = bson.M{"name": req.Name}
		cacheKey = fmt.Sprintf("name:%v", req.Name)
	}

	// get from cache, if no error and category found, direct return
	category, err := c.getCache(ctx, fmt.Sprintf(entity.CacheKeyCategory, cacheKey))
	if err == nil && !category.Id.IsZero() {
		c.logger.Info(ctx, fmt.Sprintf("cache for %v found", cacheKey))
		return category, nil
	} else if err != nil {
		c.logger.Error(ctx, err)
	}
	c.logger.Info(ctx, fmt.Sprintf("cache for %v no found", cacheKey))

	if err := c.collection.FindOne(ctx, filter).Decode(&category); err != nil && err == mongo.ErrNoDocuments {
		return category, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return category, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	// set category cache if result found from mongo
	if err := c.setCache(ctx, fmt.Sprintf(entity.CacheKeyCategory, cacheKey), category); err != nil {
		c.logger.Error(ctx, fmt.Sprintf("cache for category:%v failed to be set", req.Id.Hex()))
	}

	return category, nil
}

// Create creates new data
func (c *category) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	res, err := c.collection.InsertOne(ctx, category)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return category, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return category, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newcategory, err := c.Get(ctx, entity.Category{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newcategory, err
	}

	return newcategory, nil
}

// Update updates existing data
func (c *category) Update(ctx context.Context, category entity.Category) (entity.Category, error) {
	filter := bson.M{"_id": category.Id}

	// Prepare the update document with non-zero values only
	update := bson.M{"$set": bson.M{}}

	// Helper function to add a field to the update document if it has a non-zero value
	addField := func(fieldName string, fieldValue interface{}) {
		if !reflect.ValueOf(fieldValue).IsZero() {
			update["$set"].(bson.M)[fieldName] = fieldValue
		}
	}

	addField("name", category.Name)

	// If no non-zero fields, return the original category without updating
	if len(update["$set"].(bson.M)) == 0 {
		return c.Get(ctx, entity.Category{Id: category.Id})
	}

	_, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return category, errors.NewWithCode(codes.CodeNoSQLUpdate, err.Error())
	}

	newcategory, err := c.Get(ctx, entity.Category{Id: category.Id})
	if err != nil {
		return newcategory, err
	}

	return newcategory, nil
}

// Delete deletes existing data
func (c *category) Delete(ctx context.Context, category entity.Category) error {
	filter := bson.M{"_id": category.Id}

	res, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewWithCode(codes.CodeNoSQLDelete, err.Error())
	}

	if res.DeletedCount < 1 {
		return errors.NewWithCode(codes.CodeNoSQLNoRowsAffected, fmt.Sprintf("failed to delete id %v", category.Id.Hex()))
	}

	return nil
}
