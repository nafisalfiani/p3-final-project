package user

import (
	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func fromProto(in *User) entity.User {
	id, _ := primitive.ObjectIDFromHex(in.GetId())
	return entity.User{
		Id:       id,
		Name:     in.GetName(),
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
}

func toProto(user entity.User) *User {
	return &User{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		CreatedBy: user.CreatedBy,
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		UpdatedBy: user.UpdatedBy,
	}
}
