package user

import (
	"context"
	"fmt"
	"reflect"

	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
	cache      cache.Interface
	broker     broker.Interface
}

type Interface interface {
	List(ctx context.Context) ([]entity.User, error)
	Get(ctx context.Context, filter entity.User) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
}

// Init creates user domain
func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection, cache cache.Interface, broker broker.Interface) Interface {
	return &user{
		logger:     logger,
		json:       json,
		collection: db,
		cache:      cache,
		broker:     broker,
	}
}

// List returns list of users
func (u *user) List(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}
	cursor, err := u.collection.Find(ctx, bson.D{})
	if err != nil {
		return users, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return users, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return users, nil
}

// Get returns specific user by email
func (u *user) Get(ctx context.Context, req entity.User) (entity.User, error) {
	var user entity.User
	var filter any
	var redisKey string

	u.logger.Debug(ctx, req)
	switch {
	case req.Username != "":
		filter = bson.M{"username": req.Username}
		redisKey = fmt.Sprintf("username:%v", req.Username)
	case req.Email != "":
		filter = bson.M{"email": req.Email}
		redisKey = fmt.Sprintf("email:%v", req.Email)
	case req.Name != "":
		filter = bson.M{"name": req.Name}
		redisKey = fmt.Sprintf("name:%v", req.Name)
	case req.Id.String() != "":
		filter = bson.M{"_id": req.Id}
		redisKey = fmt.Sprintf("id:%v", req.Id.Hex())
	}

	// get from cache, if no error and user found, direct return
	user, err := u.getCache(ctx, fmt.Sprintf(entity.RedisKeyUser, redisKey))
	if err == nil && !user.Id.IsZero() {
		u.logger.Info(ctx, fmt.Sprintf("cache for %v found", redisKey))
		return user, nil
	} else if err != nil {
		u.logger.Error(ctx, err)
	}
	u.logger.Info(ctx, fmt.Sprintf("cache for %v no found", redisKey))

	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil && err == mongo.ErrNoDocuments {
		return user, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return user, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	// set user cache if result found from mongo
	if err := u.setCache(ctx, redisKey, user); err != nil {
		u.logger.Error(ctx, fmt.Sprintf("cache for user:%v failed to be set", req.Id.Hex()))
	}

	return user, nil
}

// Create creates new data
func (u *user) Create(ctx context.Context, user entity.User) (entity.User, error) {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return user, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return user, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newUser, err := u.Get(ctx, entity.User{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Update updates existing data
func (u *user) Update(ctx context.Context, user entity.User) (entity.User, error) {
	filter := bson.M{"_id": user.Id}

	// Prepare the update document with non-zero values only
	update := bson.M{"$set": bson.M{}}

	// Helper function to add a field to the update document if it has a non-zero value
	addField := func(fieldName string, fieldValue interface{}) {
		if !reflect.ValueOf(fieldValue).IsZero() {
			update["$set"].(bson.M)[fieldName] = fieldValue
		}
	}

	addField("name", user.Name)
	addField("username", user.Username)
	addField("email", user.Email)
	addField("is_email_verified", user.IsEmailVerified)
	addField("password", user.Password)
	addField("role", user.Role)
	addField("created_at", user.CreatedAt)
	addField("created_by", user.CreatedBy)
	addField("updated_at", user.UpdatedAt)
	addField("updated_by", user.UpdatedBy)

	// If no non-zero fields, return the original user without updating
	if len(update["$set"].(bson.M)) == 0 {
		return u.Get(ctx, entity.User{Id: user.Id})
	}

	_, err := u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return user, errors.NewWithCode(codes.CodeNoSQLUpdate, err.Error())
	}

	newUser, err := u.Get(ctx, entity.User{Id: user.Id})
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Delete deletes existing data
func (u *user) Delete(ctx context.Context, user entity.User) error {
	filter := bson.M{"_id": user.Id}

	res, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.NewWithCode(codes.CodeNoSQLDelete, err.Error())
	}

	if res.DeletedCount < 1 {
		return errors.NewWithCode(codes.CodeNoSQLNoRowsAffected, fmt.Sprintf("failed to delete id %v", user.Id.Hex()))
	}

	return nil
}
