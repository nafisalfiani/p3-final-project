package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	SellerId    string             `json:"seller_id" bson:"seller_id,omitempty"`
	StartDate   time.Time          `json:"start_date" bson:"start_date,omitempty"`
	Category    Category           `json:"category" bson:"category,omitempty"`
	Region      Region             `json:"region" bson:"region,omitempty"`
	Venue       Venue              `json:"venue" bson:"venue,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	CreatedBy   string             `json:"created_by" bson:"created_by,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by,omitempty"`
}

type Venue struct {
	Name         string     `json:"name" bson:"name,omitempty"`
	EntranceGate string     `json:"entrance_gate" bson:"entrance_gate,omitempty"`
	Coordinate   Coordinate `json:"coordinate" bson:"coordinate,omitempty"`
}

type Coordinate struct {
	Lat  string `json:"lat" bson:"lat,omitempty"`
	Long string `json:"long" bson:"long,omitempty"`
}
