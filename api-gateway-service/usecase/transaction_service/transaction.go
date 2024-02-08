package transactionservice

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/handler/grpc/wallet"
)

type Interface interface {
	Close()
}

type accountSvc struct {
	logger      log.Interface
	conn        *grpc.ClientConn
	transaction transaction.TransactionServiceClient
	wallet      wallet.WalletServiceClient
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
		logger:      logger,
		conn:        conn,
		transaction: transaction.NewTransactionServiceClient(conn),
		wallet:      wallet.NewWalletServiceClient(conn),
	}
}

func (a *accountSvc) Close() {
	err := a.conn.Close()
	if err != nil {
		a.logger.Error(context.Background(), err)
	}
}
