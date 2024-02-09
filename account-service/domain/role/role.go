package role

import (
	"context"
	"fmt"

	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/codes"
	"github.com/nafisalfiani/p3-final-project/lib/errors"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type role struct {
	logger     log.Interface
	json       parser.JSONInterface
	collection *mongo.Collection
	cache      cache.Interface
}

type Interface interface {
	List(ctx context.Context) ([]entity.Role, error)
	Get(ctx context.Context, role entity.Role) (entity.Role, error)
	Create(ctx context.Context, role entity.Role) (entity.Role, error)
}

// Init creates role domain
func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Collection, cache cache.Interface) Interface {
	return &role{
		logger:     logger,
		json:       json,
		collection: db,
		cache:      cache,
	}
}

func (r *role) List(ctx context.Context) ([]entity.Role, error) {
	roles := []entity.Role{}
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return roles, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &roles); err != nil {
		return roles, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	return roles, nil
}

func (r *role) Get(ctx context.Context, role entity.Role) (entity.Role, error) {
	var filter any
	var cacheKey string

	switch {
	case role.Code != "":
		filter = bson.M{"code": role.Code}
		cacheKey = fmt.Sprintf("code:%v", role.Code)
	case !role.Id.IsZero():
		filter = bson.M{"_id": role.Id}
		cacheKey = fmt.Sprintf("id:%v", role.Id.Hex())
	}

	role, err := r.getCache(ctx, fmt.Sprintf(entity.CacheKeyRole, cacheKey))
	if err == nil && !role.Id.IsZero() {
		r.logger.Info(ctx, fmt.Sprintf("cache for role:%v found", role.Id.Hex()))
		return role, nil
	} else if err != nil {
		r.logger.Error(ctx, err)
	}
	r.logger.Info(ctx, fmt.Sprintf("cache for role:%v not found", role.Id.Hex()))

	if err := r.collection.FindOne(ctx, filter).Decode(&role); err != nil && err == mongo.ErrNoDocuments {
		return role, errors.NewWithCode(codes.CodeNoSQLRecordDoesNotExist, err.Error())
	} else if err != nil {
		return role, errors.NewWithCode(codes.CodeNoSQLRead, err.Error())
	}

	// set role cache if result found
	if err := r.setCache(ctx, fmt.Sprintf(entity.CacheKeyRole, cacheKey), role); err != nil {
		r.logger.Error(ctx, fmt.Sprintf("cache for user:%v failed to be set", role.Id.Hex()))
	}

	return role, nil
}

func (r *role) Create(ctx context.Context, role entity.Role) (entity.Role, error) {
	res, err := r.collection.InsertOne(ctx, role)
	if err != nil && mongo.IsDuplicateKeyError(err) {
		return role, errors.NewWithCode(codes.CodeNoSQLConflict, err.Error())
	} else if err != nil {
		return role, errors.NewWithCode(codes.CodeNoSQLInsert, err.Error())
	}

	newRole, err := r.Get(ctx, entity.Role{Id: res.InsertedID.(primitive.ObjectID)})
	if err != nil {
		return newRole, err
	}

	return newRole, nil
}
