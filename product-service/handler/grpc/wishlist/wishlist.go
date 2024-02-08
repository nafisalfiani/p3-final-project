package wishlist

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/usecase/wishlist"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcWish struct {
	log      log.Interface
	wishlist wishlist.Interface
}

func Init(log log.Interface, wishlist wishlist.Interface) WishlistServiceServer {
	return &grpcWish{
		log:      log,
		wishlist: wishlist,
	}
}

func (w *grpcWish) mustEmbedUnimplementedWishlistServiceServer() {}

func (w *grpcWish) GetWishlist(ctx context.Context, in *Wishlist) (*Wishlist, error) {
	wishlist, err := w.wishlist.Get(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wishlist), nil
}

func (w *grpcWish) CreateWishlist(ctx context.Context, in *Wishlist) (*Wishlist, error) {
	wishlist, err := w.wishlist.Create(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wishlist), nil
}

func (w *grpcWish) UpdateWishlist(ctx context.Context, in *Wishlist) (*Wishlist, error) {
	wishlist, err := w.wishlist.Update(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(wishlist), nil
}

func (w *grpcWish) DeleteWishlist(ctx context.Context, in *Wishlist) (*emptypb.Empty, error) {
	if err := w.wishlist.Delete(ctx, fromProto(in)); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (w *grpcWish) GetWishlists(ctx context.Context, in *Wishlist) (*WishlistList, error) {
	wishlists, err := w.wishlist.List(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	res := &WishlistList{}
	for i := range wishlists {
		res.Wishlists = append(res.Wishlists, toProto(wishlists[i]))
	}

	return res, nil
}
