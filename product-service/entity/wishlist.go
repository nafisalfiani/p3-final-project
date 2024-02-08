package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wishlist struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Category        Category           `json:"category" bson:"category,omitempty"`
	Region          Region             `json:"region" bson:"region,omitempty"`
	SubscribedUsers []string           `json:"subscribed_users" bson:"subscribed_users,omitempty"`
}
