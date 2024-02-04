package domain

import (
	"github.com/nafisalfiani/p3-final-project/account-service/domain/role"
	"github.com/nafisalfiani/p3-final-project/account-service/domain/user"
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type Domains struct {
	User user.Interface
	Role role.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Client, cache cache.Interface, broker broker.Interface) *Domains {
	return &Domains{
		User: user.Init(logger, json, db.Database("account-service").Collection("user"), cache, broker),
		Role: role.Init(logger, json, db.Database("account-service").Collection("role"), cache),
	}
}
