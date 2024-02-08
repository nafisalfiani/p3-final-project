package domain

import (
	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/cache"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/category"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/region"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/domain/wishlist"
	"go.mongodb.org/mongo-driver/mongo"
)

type Domains struct {
	Ticket   ticket.Interface
	Category category.Interface
	Region   region.Interface
	Wishlist wishlist.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *mongo.Client, cache cache.Interface, broker broker.Interface) *Domains {
	return &Domains{
		Ticket:   ticket.Init(logger, json, db.Database("product-service").Collection("product"), cache, broker),
		Category: category.Init(logger, json, db.Database("product-service").Collection("category"), cache),
		Region:   region.Init(logger, json, db.Database("product-service").Collection("region"), cache),
		Wishlist: wishlist.Init(logger, json, db.Database("product-service").Collection("wishlist")),
	}
}
