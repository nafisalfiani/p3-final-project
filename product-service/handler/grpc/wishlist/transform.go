package wishlist

import (
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func toProto(ticket entity.Wishlist) *Wishlist {
	return &Wishlist{
		Id:              ticket.Id.Hex(),
		CategoryName:    ticket.Category.Name,
		RegionName:      ticket.Region.Name,
		SubscribedUsers: ticket.SubscribedUsers,
	}
}

func fromProto(in *Wishlist) entity.Wishlist {
	id, _ := primitive.ObjectIDFromHex(in.GetId())
	return entity.Wishlist{
		Id: id,
		Category: entity.Category{
			Name: in.GetCategoryName(),
		},
		Region: entity.Region{
			Name: in.GetRegionName(),
		},
		SubscribedUsers: in.GetSubscribedUsers(),
	}
}
