package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Region struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
