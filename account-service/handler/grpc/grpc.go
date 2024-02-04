package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-playground/validator/v10"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc/auth"
	"github.com/nafisalfiani/p3-final-project/account-service/handler/grpc/user"
	"github.com/nafisalfiani/p3-final-project/account-service/usecase"
	jwtAuth "github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/header"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Interface interface {
	Run()
}

type Config struct {
	Base string `env:"GRPC_BASE"`
	Port int    `env:"GRPC_PORT"`
}

type grpcServer struct {
	cfg    Config
	log    log.Interface
	auth   jwtAuth.Interface
	server *grpc.Server
}

func Init(cfg Config, log log.Interface, uc *usecase.Usecases, sec security.Interface, jwtAuth jwtAuth.Interface, validator *validator.Validate) Interface {
	srv := &grpcServer{
		cfg:  cfg,
		log:  log,
		auth: jwtAuth,
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(srv.authFunc)),
	)

	user.RegisterUserServiceServer(s, user.Init(log, uc.User, jwtAuth, validator))
	auth.RegisterAuthServiceServer(s, auth.Init(log, sec, jwtAuth, uc.User, uc.Role))

	reflection.Register(s)

	srv.server = s

	return srv
}

func (g *grpcServer) Run() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", g.cfg.Base, g.cfg.Port))
	if err != nil {
		g.log.Fatal(ctx, err)
	}

	g.log.Info(ctx, fmt.Sprintf("Listening and Serving GRPC on %v:%v", g.cfg.Base, g.cfg.Port))
	if err := g.server.Serve(listener); err != nil {
		g.log.Fatal(ctx, err)
	}
}

// wrapper to connect to grpc package
func Dial(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(target, opts...)
}

// wrapper to connect to grpc package
func WithInsecure() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}

func (g *grpcServer) authFunc(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "failed to get method name")
	}
	g.log.Info(ctx, fmt.Sprintf("authenticating method %v", method))

	// Skip authentication for specific methods
	if method == "/auth.AuthService/Register" || method == "/auth.AuthService/Login" || method == "/user.UserService/VerifyUserEmail" {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not provided")
	}

	user, err := g.auth.VerifyToken(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return g.auth.SetUserAuthInfo(ctx, user, &jwtAuth.Token{TokenType: header.AuthorizationBearer, AccessToken: token}), nil
}
