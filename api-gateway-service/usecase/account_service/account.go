package accountservice

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc/auth"
	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc/user"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/log"
)

type Interface interface {
	Close()
	SignIn(ctx context.Context, req entity.LoginRequest) (entity.LoginResponse, error)
	Register(ctx context.Context, req entity.RegisterRequest) (entity.RegisterResponse, error)
}

type accountSvc struct {
	logger log.Interface
	conn   *grpc.ClientConn
	auth   auth.AuthServiceClient
	user   user.UserServiceClient
}

type Config struct {
	Base string
	Port int
}

func Init(cfg Config, logger log.Interface) Interface {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", cfg.Base, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(context.Background(), err)
	}

	return &accountSvc{
		logger: logger,
		conn:   conn,
		auth:   auth.NewAuthServiceClient(conn),
		user:   user.NewUserServiceClient(conn),
	}
}

func (a *accountSvc) Close() {
	err := a.conn.Close()
	if err != nil {
		a.logger.Error(context.Background(), err)
	}
}
