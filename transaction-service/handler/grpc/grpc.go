package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-playground/validator/v10"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	jwtAuth "github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/header"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/security"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/wallet"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase"
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

	transaction.RegisterTransactionServiceServer(s, transaction.Init(log, uc.Transaction))
	wallet.RegisterWalletServiceServer(s, wallet.Init(log, uc.Wallet))

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
