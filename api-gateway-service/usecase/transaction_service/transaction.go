package transactionservice

import (
	"context"

	"github.com/nafisalfiani/p3-final-project/api-gateway-service/entity"
)

func (t *trxSvc) ListTransaction(ctx context.Context, in entity.Transaction) ([]entity.Transaction, error) {
	transactions, err := t.transaction.GetTransactions(ctx, toTransactionProto(in))
	if err != nil {
		return nil, err
	}

	var res []entity.Transaction
	for _, v := range transactions.GetTransactions() {
		res = append(res, fromTransactionProto(v))
	}

	return res, nil
}

func (t *trxSvc) CreateTransaction(ctx context.Context, in entity.Transaction) (entity.Transaction, error) {
	transaction, err := t.transaction.CreateTransaction(ctx, toTransactionProto(in))
	if err != nil {
		return fromTransactionProto(transaction), err
	}

	return fromTransactionProto(transaction), nil
}

func (t *trxSvc) UpdateTransaction(ctx context.Context, in entity.Transaction) (entity.Transaction, error) {
	transaction, err := t.transaction.UpdateTransaction(ctx, toTransactionProto(in))
	if err != nil {
		return fromTransactionProto(transaction), err
	}

	return fromTransactionProto(transaction), nil
}
