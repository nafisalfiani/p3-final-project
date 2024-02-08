package entity

import (
	"time"

	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	Username        string             `json:"username" bson:"username,omitempty"`
	Email           string             `json:"email" bson:"email,omitempty"`
	IsEmailVerified bool               `json:"-" bson:"is_email_verified,omitempty"`
	Password        string             `json:"password" bson:"password,omitempty"`
	Role            Role               `json:"role" bson:"role,omitempty"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at,omitempty"`
	CreatedBy       string             `json:"created_by" bson:"created_by,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	UpdatedBy       string             `json:"updated_by" bson:"updated_by,omitempty"`
}

func (u User) ToAuthUser() auth.User {
	user := auth.User{
		Id:       u.Id.Hex(),
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role.Code,
		Scopes:   u.Role.Scopes,
	}

	return user
}
