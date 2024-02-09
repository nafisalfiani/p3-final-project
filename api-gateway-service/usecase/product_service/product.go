package productservice

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/category"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/region"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/wishlist"
)

type Interface interface {
	Close()

	// tickets
	ListCategory(ctx context.Context) ([]entity.Category, error)
	ListRegion(ctx context.Context) ([]entity.Region, error)
	ListTicket(ctx context.Context, in entity.Ticket) ([]entity.Ticket, error)
	GetTicket(ctx context.Context, in entity.Ticket) (entity.Ticket, error)
	RegisterTicketForSale(ctx context.Context, in entity.Ticket) (entity.Ticket, error)
	UpdateTicketInfo(ctx context.Context, in entity.Ticket) (entity.Ticket, error)
	TakeDownTicket(ctx context.Context, in entity.Ticket) error

	// wishlists
	ListWishlist(ctx context.Context, wishlist entity.Wishlist) ([]entity.Wishlist, error)
	GetWishlistSubscriber(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Subscribe(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Unsubscribe(ctx context.Context, wishlist entity.Wishlist) error
}

type productSvc struct {
	logger   log.Interface
	auth     auth.Interface
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

func Init(cfg Config, logger log.Interface, auth auth.Interface) Interface {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", cfg.Base, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(context.Background(), err)
	}

	return &productSvc{
		logger:   logger,
		conn:     conn,
		auth:     auth,
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
