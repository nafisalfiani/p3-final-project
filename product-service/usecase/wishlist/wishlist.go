package wishlist

import (
	"context"

	wishlistDom "github.com/nafisalfiani/p3-final-project/product-service/domain/wishlist"
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
)

type Interface interface {
	List(ctx context.Context, filter entity.Wishlist) ([]entity.Wishlist, error)
	Get(ctx context.Context, filter entity.Wishlist) (entity.Wishlist, error)
	Create(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Update(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error)
	Delete(ctx context.Context, wishlist entity.Wishlist) error
}

type wishlist struct {
	wishlist wishlistDom.Interface
}

func Init(prd wishlistDom.Interface) Interface {
	return &wishlist{
		wishlist: prd,
	}
}

func (w *wishlist) List(ctx context.Context, filter entity.Wishlist) ([]entity.Wishlist, error) {
	return w.wishlist.List(ctx, filter)
}

func (w *wishlist) Get(ctx context.Context, filter entity.Wishlist) (entity.Wishlist, error) {
	return w.wishlist.Get(ctx, filter)
}

func (w *wishlist) Create(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error) {
	return w.wishlist.Create(ctx, wishlist)
}

func (w *wishlist) Update(ctx context.Context, wishlist entity.Wishlist) (entity.Wishlist, error) {
	return w.wishlist.Update(ctx, wishlist)
}

func (w *wishlist) Delete(ctx context.Context, wishlist entity.Wishlist) error {
	return w.wishlist.Delete(ctx, wishlist)
}
