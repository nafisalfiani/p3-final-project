package ticket

import (
	"github.com/nafisalfiani/p3-final-project/product-service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProto(ticket entity.Ticket) *Ticket {
	return &Ticket{
		Id:                  ticket.Id.Hex(),
		Title:               ticket.Title,
		Description:         ticket.Description,
		SellerId:            ticket.SellerId,
		StartDate:           timestamppb.New(ticket.StartDate),
		CategoryName:        ticket.Category.Name,
		RegionName:          ticket.Region.Name,
		VenueName:           ticket.Venue.Name,
		VenueEntranceGate:   ticket.Venue.EntranceGate,
		VenueCoordinateLat:  ticket.Venue.Coordinate.Lat,
		VenueCoordinateLong: ticket.Venue.Coordinate.Long,
		CreatedAt:           timestamppb.New(ticket.CreatedAt),
		CreatedBy:           ticket.CreatedBy,
		UpdatedAt:           timestamppb.New(ticket.UpdatedAt),
		UpdatedBy:           ticket.UpdatedBy,
	}
}

func fromProto(in *Ticket) entity.Ticket {
	id, _ := primitive.ObjectIDFromHex(in.GetId())
	return entity.Ticket{
		Id:          id,
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		SellerId:    in.GetSellerId(),
		StartDate:   in.GetStartDate().AsTime(),
		Category: entity.Category{
			Name: in.GetCategoryName(),
		},
		Region: entity.Region{
			Name: in.GetRegionName(),
		},
		Venue: entity.Venue{
			Name:         in.GetVenueName(),
			EntranceGate: in.GetVenueEntranceGate(),
			Coordinate: entity.Coordinate{
				Lat:  in.GetVenueCoordinateLat(),
				Long: in.GetVenueCoordinateLong(),
			},
		},
		CreatedAt: in.GetCreatedAt().AsTime(),
		CreatedBy: in.GetCreatedBy(),
		UpdatedAt: in.GetUpdatedAt().AsTime(),
		UpdatedBy: in.GetUpdatedBy(),
	}
}
