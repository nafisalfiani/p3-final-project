package transactionservice

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/product-service/handler/grpc/ticket"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/wallet"
)

type Interface interface {
	Close()

	// transaction
	ListTransaction(ctx context.Context, in entity.Transaction) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, in entity.Transaction) (entity.Transaction, error)
	UpdateTransaction(ctx context.Context, in entity.Transaction) (entity.Transaction, error)

	// wallet
	GetWallet(ctx context.Context, in entity.Wallet) (entity.Wallet, error)
}

type trxSvc struct {
	logger      log.Interface
	conn        *grpc.ClientConn
	transaction transaction.TransactionServiceClient
	wallet      wallet.WalletServiceClient
	ticket      ticket.TicketServiceClient
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

	return &trxSvc{
		logger:      logger,
		conn:        conn,
		transaction: transaction.NewTransactionServiceClient(conn),
		wallet:      wallet.NewWalletServiceClient(conn),
		ticket:      ticket.NewTicketServiceClient(conn),
	}
}

func (a *trxSvc) Close() {
	err := a.conn.Close()
	if err != nil {
		a.logger.Error(context.Background(), err)
	}
}
