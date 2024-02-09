package productservice

import (
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/ticket"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/wishlist"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toTicketProto(in entity.Ticket) *ticket.Ticket {
	return &ticket.Ticket{
		Id:                  in.Id,
		Title:               in.Title,
		Description:         in.Description,
		SellerId:            in.SellerId,
		StartDate:           timestamppb.New(in.StartDate),
		SellingPrice:        in.SellingPrice,
		CategoryName:        in.Category.Name,
		RegionName:          in.Region.Name,
		VenueName:           in.Venue.Name,
		VenueEntranceGate:   in.Venue.EntranceGate,
		VenueCoordinateLat:  in.Venue.Coordinate.Lat,
		VenueCoordinateLong: in.Venue.Coordinate.Long,
		CreatedAt:           timestamppb.New(in.CreatedAt),
		CreatedBy:           in.CreatedBy,
		UpdatedAt:           timestamppb.New(in.UpdatedAt),
		UpdatedBy:           in.UpdatedBy,
	}
}

func fromTicketProto(v *ticket.Ticket) entity.Ticket {
	return entity.Ticket{
		Id:           v.GetId(),
		Title:        v.GetTitle(),
		Description:  v.GetDescription(),
		SellerId:     v.GetSellerId(),
		StartDate:    v.GetStartDate().AsTime(),
		SellingPrice: v.GetSellingPrice(),
		Category: entity.Category{
			Name: v.GetCategoryName(),
		},
		Region: entity.Region{
			Name: v.GetRegionName(),
		},
		Venue: entity.Venue{
			Name:         v.GetVenueName(),
			EntranceGate: v.GetVenueEntranceGate(),
			Coordinate: entity.Coordinate{
				Lat:  v.GetVenueCoordinateLat(),
				Long: v.GetVenueCoordinateLong(),
			},
		},
		CreatedAt: v.GetCreatedAt().AsTime(),
		CreatedBy: v.GetCreatedBy(),
		UpdatedAt: v.GetUpdatedAt().AsTime(),
		UpdatedBy: v.GetUpdatedBy(),
	}
}

func toWishlistProto(in entity.Wishlist) *wishlist.Wishlist {
	return &wishlist.Wishlist{
		Id:              in.Id,
		CategoryName:    in.CategoryName,
		RegionName:      in.RegionName,
		SubscribedUsers: in.SubscribedUsers,
	}
}

func fromWishlistProto(v *wishlist.Wishlist) entity.Wishlist {
	return entity.Wishlist{
		Id:              v.GetId(),
		CategoryName:    v.GetCategoryName(),
		RegionName:      v.GetRegionName(),
		SubscribedUsers: v.GetSubscribedUsers(),
	}
}
