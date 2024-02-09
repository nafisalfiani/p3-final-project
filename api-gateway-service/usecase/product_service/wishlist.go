package productservice

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
)

func (p *productSvc) ListWishlist(ctx context.Context, wishlist entity.Wishlist) ([]entity.Wishlist, error) {
	wishlists, err := p.wishlist.GetWishlists(ctx, toWishlistProto(wishlist))
	if err != nil {
		return nil, err
	}

	var res []entity.Wishlist
	for _, v := range wishlists.GetWishlists() {
		res = append(res, entity.Wishlist{
			Id:           v.GetId(),
			CategoryName: v.GetCategoryName(),
			RegionName:   v.GetRegionName(),
		})
	}

	return res, nil
}

func (p *productSvc) GetWishlistSubscriber(ctx context.Context, in entity.Wishlist) (entity.Wishlist, error) {
	wishlist, err := p.wishlist.GetWishlist(ctx, toWishlistProto(in))
	if err != nil {
		return entity.Wishlist{}, err
	}

	return fromWishlistProto(wishlist), nil
}

func (p *productSvc) Subscribe(ctx context.Context, in entity.Wishlist) (entity.Wishlist, error) {
	wishlist, err := p.wishlist.CreateWishlist(ctx, toWishlistProto(in))
	if err != nil {
		return entity.Wishlist{}, err
	}

	return fromWishlistProto(wishlist), nil
}

func (p *productSvc) Unsubscribe(ctx context.Context, in entity.Wishlist) error {
	wishlist, err := p.wishlist.UpdateWishlist(ctx, toWishlistProto(in))
	if err != nil {
		return err
	}
	p.logger.Debug(ctx, wishlist)

	return nil
}
