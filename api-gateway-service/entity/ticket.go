package entity

import (
	"time"
)

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Region struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Ticket struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	SellerId     string    `json:"seller_id"`
	StartDate    time.Time `json:"start_date"`
	SellingPrice float32   `json:"selling_price"`
	Category     Category  `json:"category"`
	Region       Region    `json:"region"`
	Venue        Venue     `json:"venue"`
	Status       string    `json:"status"`
	BuyerId      string    `json:"buyer_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}

type TicketCreateRequest struct {
	Title             string    `json:"title" validate:"required"`
	Description       string    `json:"description" validate:"required"`
	StartDate         time.Time `json:"start_date" validate:"required"`
	SellingPrice      float32   `json:"selling_price" validate:"required"`
	CategoryName      string    `json:"category" validate:"required"`
	RegionName        string    `json:"region" validate:"required"`
	VenueName         string    `json:"venue_name" validate:"required"`
	VenueEntranceGate string    `json:"venue_entrance_gate" validate:"required"`
}

func (t TicketCreateRequest) ToTicket() Ticket {
	return Ticket{
		Title:        t.Title,
		Description:  t.Description,
		StartDate:    t.StartDate,
		SellingPrice: t.SellingPrice,
		Category: Category{
			Name: t.CategoryName,
		},
		Region: Region{
			Name: t.RegionName,
		},
		Venue: Venue{
			Name:         t.VenueName,
			EntranceGate: t.VenueEntranceGate,
		},
	}
}

type TicketUpdateRequest struct {
	Id           string    `uri:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	SellingPrice float32   `json:"selling_price"`
}

type TicketGetRequest struct {
	Id string `uri:"id"`
}

func (t TicketUpdateRequest) ToTicket() Ticket {
	return Ticket{
		Id:           t.Id,
		Title:        t.Title,
		Description:  t.Description,
		StartDate:    t.StartDate,
		SellingPrice: t.SellingPrice,
	}
}

type Venue struct {
	Name         string     `json:"name"`
	EntranceGate string     `json:"entrance_gate"`
	Coordinate   Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}
