package wishlist

import (
	"context"
	"fmt"
	"reflect"

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
	List(ctx context.Context, filter entity.Wishlist) ([]entity.Wishlist, error)
	Get(ctx context.Context, filter entity.Wishlist) (entity.Wishlist, error)
	Create(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Update(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Delete(ctx context.Context, wishlist entity.Wishlist) error
}

type wishlist struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection) Interface {
	return &wishlist{
		logger:     logger,
		json:       json,
		collection: db,
	}
}

// List returns list of wishlists
func (w *wishlist) List(ctx context.Context, req entity.Wishlist) ([]entity.Wishlist, error) {
	wishlists := []entity.Wishlist{}
	var filter any

	switch {
	case len(req.SubscribedUsers) > 0:
		filter = bson.M{"subscribed_user": bson.M{
			"$in": req.SubscribedUsers,
		}}
	case req.Category.Name != "":
		filter = bson.M{"category.name": req.Category.Name}
	case req.Region.Name != "":
		filter = bson.M{"region.name": req.Region.Name}
	}

	cursor, err := w.collection.Find(ctx, filter)
	if err != nil {
		return wishlists, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &wishlists); err != nil {
		return wishlists, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return wishlists, nil
}

// Get returns specific wishlist
func (w *wishlist) Get(ctx context.Context, req entity.Wishlist) (entity.Wishlist, error) {
	var wishlist entity.Wishlist
	var filter any

	w.logger.Debug(ctx, req)
	switch {
	case !req.Id.IsZero():
		filter = bson.M{"_id": req.Id}
	case req.Category.Name != "" && req.Region.Name != "":
		filter = bson.M{
			"region.name":   req.Region.Name,
			"category.name": req.Category.Name,
		}
	}

	if err := w.collection.FindOne(ctx, filter).Decode(&wishlist); err != nil && err == mongo.ErrNoDocuments {
		return wishlist, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return wishlist, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return wishlist, nil
}

// Create creates new data
func (w *wishlist) Create(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error) {
	res, err := w.collection.InsertOne(ctx, wishlist)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return wishlist, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return wishlist, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newwishlist, err := w.Get(ctx, entity.Wishlist{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newwishlist, err
	}

	return newwishlist, nil
}

// Update updates existing data
func (w *wishlist) Update(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error) {
	filter := bson.M{"_id": wishlist.Id}

	// Prepare the update document with non-zero values only
	update := bson.M{"$set": bson.M{}}

	// Helper function to add a field to the update document if it has a non-zero value
	addField := func(fieldName string, fieldValue interface{}) {
		if !reflect.ValueOf(fieldValue).IsZero() {
			update["$set"].(bson.M)[fieldName] = fieldValue
		}
	}

	addField("category", wishlist.Category)
	addField("region", wishlist.Region)
	addField("subscribed_users", wishlist.SubscribedUsers)

	// If no non-zero fields, return the original wishlist without updating
	if len(update["$set"].(bson.M)) == 0 {
		return w.Get(ctx, entity.Wishlist{Id: wishlist.Id})
	}

	_, err := w.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return wishlist, errors.NewWithCode(codes.CodeNoSQLUpdate, err.Error())
	}

	newwishlist, err := w.Get(ctx, entity.Wishlist{Id: wishlist.Id})
	if err != nil {
		return newwishlist, err
	}

	return newwishlist, nil
}

// Delete deletes existing data
func (w *wishlist) Delete(ctx context.Context, wishlist entity.Wishlist) error {
	filter := bson.M{"_id": wishlist.Id}

	res, err := w.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewWithCode(codes.CodeNoSQLDelete, err.Error())
	}

	if res.DeletedCount < 1 {
		return errors.NewWithCode(codes.CodeNoSQLNoRowsAffected, fmt.Sprintf("failed to delete id %v", wishlist.Id.Hex()))
	}

	return nil
}
