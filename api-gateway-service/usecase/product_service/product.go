package productservice

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/category"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/region"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/wishlist"
)

type Interface interface {
	Close()
}

type productSvc struct {
	logger   log.Interface
	conn     *grpc.ClientConn
	ticket   ticket.TicketServiceClient
	category category.CategoryServiceClient
	region   region.RegionServiceClient
	wishlist wishlist.WishlistServiceClient
}

type Config struct {
	Base string
	Port int
}

func Init(cfg Config, logger log.Interface) Interface {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", cfg.Base, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(context.Background(), err)
	}

	return &productSvc{
		logger:   logger,
		conn:     conn,
		ticket:   ticket.NewTicketServiceClient(conn),
		category: category.NewCategoryServiceClient(conn),
		region:   region.NewRegionServiceClient(conn),
		wishlist: wishlist.NewWishlistServiceClient(conn),
	}
}

func (a *productSvc) Close() {
	err := a.conn.Close()
	if err != nil {
		a.logger.Error(context.Background(), err)
	}
}
