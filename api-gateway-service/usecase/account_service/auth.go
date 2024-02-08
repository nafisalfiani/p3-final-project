package accountservice

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc/auth"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
)

func (a *accountSvc) SignIn(ctx context.Context, req entity.LoginRequest) (entity.LoginResponse, error) {
	res, err := a.auth.Login(ctx, &auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return entity.LoginResponse{}, err
	}

	return entity.LoginResponse{
		TokenType:       res.GetTokenType(),
		AccessToken:     res.GetAccessToken(),
		AccessExpiresIn: res.GetAccessExpiresIn(),
	}, nil
}

func (a *accountSvc) Register(ctx context.Context, req entity.RegisterRequest) (entity.RegisterResponse, error) {
	res, err := a.auth.Register(ctx, &auth.RegisterRequest{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return entity.RegisterResponse{}, err
	}

	return entity.RegisterResponse{
		Id:       res.GetId(),
		Username: res.GetUsername(),
		Name:     res.GetName(),
		Email:    res.GetEmail(),
		Role:     res.GetRole(),
	}, nil
}
