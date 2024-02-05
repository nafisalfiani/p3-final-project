package user

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/p3-final-project/account-service/entity"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase/user"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userGrpcServer struct {
	log  log.Interface
	user user.Interface
	auth auth.Interface
}

func Init(log log.Interface, user user.Interface, auth auth.Interface, validator *validator.Validate) *userGrpcServer {
	return &userGrpcServer{
		log:  log,
		user: user,
		auth: auth,
	}
}

func (u *userGrpcServer) mustEmbedUnimplementedUserServiceServer() {}

func (u *userGrpcServer) GetUser(ctx context.Context, req *User) (*User, error) {
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		u.log.Error(ctx, err)
	}

	user, err := u.user.Get(ctx, entity.User{
		Id:    id,
		Name:  req.GetName(),
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	res := &User{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		CreatedBy: user.CreatedBy,
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		UpdatedBy: user.UpdatedBy,
	}

	return res, nil
}

func (u *userGrpcServer) CreateUser(ctx context.Context, req *User) (*User, error) {
	newUser, err := u.user.Create(ctx, entity.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}

	res := &User{
		Id:    newUser.Id.Hex(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return res, nil
}

func (u *userGrpcServer) UpdateUser(ctx context.Context, req *User) (*User, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := u.user.Update(ctx, entity.User{
		Id:       id,
		Name:     req.GetName(),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	u.log.Debug(ctx, user)

	res := &User{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		CreatedBy: user.CreatedBy,
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		UpdatedBy: user.UpdatedBy,
	}

	return res, nil
}

func (u *userGrpcServer) DeleteUser(ctx context.Context, req *User) (*emptypb.Empty, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	if err := u.user.Delete(ctx, entity.User{
		Id: id,
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *userGrpcServer) GetUsers(ctx context.Context, in *emptypb.Empty) (*UserList, error) {
	users, err := u.user.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &UserList{}
	for i := range users {
		res.Users = append(res.Users, &User{
			Id:        users[i].Id.Hex(),
			Name:      users[i].Name,
			Email:     users[i].Email,
			Username:  users[i].Username,
			Password:  users[i].Password,
			CreatedAt: timestamppb.New(users[i].CreatedAt),
			CreatedBy: users[i].CreatedBy,
			UpdatedAt: timestamppb.New(users[i].UpdatedAt),
			UpdatedBy: users[i].UpdatedBy,
		})
	}

	return res, nil
}

func (u *userGrpcServer) VerifyUserEmail(ctx context.Context, in *User) (*User, error) {
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.user.Update(ctx, entity.User{
		Id:              id,
		IsEmailVerified: true,
		UpdatedAt:       time.Now(),
		UpdatedBy:       id.Hex(),
	})
	if err != nil {
		return nil, err
	}

	res := &User{
		Id:        updatedUser.Id.Hex(),
		Name:      updatedUser.Name,
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		CreatedAt: timestamppb.New(updatedUser.CreatedAt),
		CreatedBy: updatedUser.CreatedBy,
		UpdatedAt: timestamppb.New(updatedUser.UpdatedAt),
		UpdatedBy: updatedUser.UpdatedBy,
	}

	return res, nil
}
