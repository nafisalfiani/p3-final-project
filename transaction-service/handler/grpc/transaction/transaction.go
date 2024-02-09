package transaction

import (
	context "context"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/usecase/transaction"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcTrx struct {
	log log.Interface
	trx transaction.Interface
}

func Init(log log.Interface, trx transaction.Interface) TransactionServiceServer {
	return &grpcTrx{
		log: log,
		trx: trx,
	}
}

func (t *grpcTrx) mustEmbedUnimplementedTransactionServiceServer() {}

func (t *grpcTrx) GetTransaction(ctx context.Context, in *Transaction) (*Transaction, error) {
	transaction, err := t.trx.Get(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(transaction), nil
}

func (t *grpcTrx) CreateTransaction(ctx context.Context, in *Transaction) (*Transaction, error) {
	transaction, err := t.trx.Create(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(transaction), nil
}

func (t *grpcTrx) UpdateTransaction(ctx context.Context, in *Transaction) (*Transaction, error) {
	transaction, err := t.trx.Update(ctx, fromProto(in))
	if err != nil {
		return nil, err
	}

	return toProto(transaction), nil
}

func (t *grpcTrx) DeleteTransaction(ctx context.Context, in *Transaction) (*emptypb.Empty, error) {
	if err := t.trx.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (t *grpcTrx) GetTransactions(ctx context.Context, in *Transaction) (*TransactionList, error) {
	transactions, err := t.trx.List(ctx)
	if err != nil {
		return nil, err
	}

	res := &TransactionList{}
	for i := range transactions {
		res.Transactions = append(res.Transactions, toProto(transactions[i]))
	}

	return res, nil
}
