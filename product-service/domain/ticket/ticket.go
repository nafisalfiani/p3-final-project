package ticket

import (
	"context"
	"fmt"
	"reflect"

	"github.com/nafisalfiani/p3-final-project/lib/broker"
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
	List(ctx context.Context) ([]entity.Ticket, error)
	Get(ctx context.Context, filter entity.Ticket) (entity.Ticket, error)
	Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error)
	Delete(ctx context.Context, ticket entity.Ticket) error
}

type ticket struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
	cache      cache.Interface
	broker     broker.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection, cache cache.Interface, broker broker.Interface) Interface {
	return &ticket{
		logger:     logger,
		json:       json,
		collection: db,
		cache:      cache,
		broker:     broker,
	}
}

// List returns list of tickets
func (t *ticket) List(ctx context.Context) ([]entity.Ticket, error) {
	tickets := []entity.Ticket{}
	cursor, err := t.collection.Find(ctx, bson.M{})
	if err != nil {
		return tickets, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tickets); err != nil {
		return tickets, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return tickets, nil
}

// Get returns specific ticket
func (t *ticket) Get(ctx context.Context, req entity.Ticket) (entity.Ticket, error) {
	var ticket entity.Ticket
	var filter any
	var cacheKey string

	t.logger.Debug(ctx, req)
	switch {
	case !req.Id.IsZero():
		filter = bson.M{"_id": req.Id}
		cacheKey = fmt.Sprintf("id:%v", req.Id)
	case req.Title != "":
		filter = bson.M{"title": req.Title}
		cacheKey = fmt.Sprintf("title:%v", req.Title)
	}

	// get from cache, if no error and ticket found, direct return
	ticket, err := t.getCache(ctx, fmt.Sprintf(entity.CacheKeyTicket, cacheKey))
	if err == nil && !ticket.Id.IsZero() {
		t.logger.Info(ctx, fmt.Sprintf("cache for %v found", cacheKey))
		return ticket, nil
	} else if err != nil {
		t.logger.Error(ctx, err)
	}
	t.logger.Info(ctx, fmt.Sprintf("cache for %v no found", cacheKey))

	if err := t.collection.FindOne(ctx, filter).Decode(&ticket); err != nil && err == mongo.ErrNoDocuments {
		return ticket, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return ticket, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	// set ticket cache if result found from mongo
	if err := t.setCache(ctx, cacheKey, ticket); err != nil {
		t.logger.Error(ctx, fmt.Sprintf("cache for ticket:%v failed to be set", req.Id.Hex()))
	}

	return ticket, nil
}

// Create creates new data
func (t *ticket) Create(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	res, err := t.collection.InsertOne(ctx, ticket)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return ticket, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return ticket, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newticket, err := t.Get(ctx, entity.Ticket{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newticket, err
	}

	return newticket, nil
}

// Update updates existing data
func (t *ticket) Update(ctx context.Context, ticket entity.Ticket) (entity.Ticket, error) {
	filter := bson.M{"_id": ticket.Id}

	// Prepare the update document with non-zero values only
	update := bson.M{"$set": bson.M{}}

	// Helper function to add a field to the update document if it has a non-zero value
	addField := func(fieldName string, fieldValue interface{}) {
		if !reflect.ValueOf(fieldValue).IsZero() {
			update["$set"].(bson.M)[fieldName] = fieldValue
		}
	}

	addField("title", ticket.Title)
	addField("description", ticket.Description)

	// If no non-zero fields, return the original ticket without updating
	if len(update["$set"].(bson.M)) == 0 {
		return t.Get(ctx, entity.Ticket{Id: ticket.Id})
	}

	_, err := t.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ticket, errors.NewWithCode(codes.CodeNoSQLUpdate, err.Error())
	}

	newticket, err := t.Get(ctx, entity.Ticket{Id: ticket.Id})
	if err != nil {
		return newticket, err
	}

	return newticket, nil
}

// Delete deletes existing data
func (t *ticket) Delete(ctx context.Context, ticket entity.Ticket) error {
	filter := bson.M{"_id": ticket.Id}

	res, err := t.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewWithCode(codes.CodeNoSQLDelete, err.Error())
	}

	if res.DeletedCount < 1 {
		return errors.NewWithCode(codes.CodeNoSQLNoRowsAffected, fmt.Sprintf("failed to delete id %v", ticket.Id.Hex()))
	}

	return nil
}
