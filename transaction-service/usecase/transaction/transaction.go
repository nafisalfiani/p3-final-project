package transaction

import (
	"context"

	transactionDom "github.com/nafisalfiani/p3-final-project/transaction-service/domain/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
)

type Interface interface {
	List(ctx context.Context) ([]entity.Transaction, error)
	Get(ctx context.Context, filter entity.Transaction) (entity.Transaction, error)
	Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	Delete(ctx context.Context, transactionId string) error
}

type transaction struct {
	transaction transactionDom.Interface
}

func Init(prd transactionDom.Interface) Interface {
	return &transaction{
		transaction: prd,
	}
}

func (c *transaction) List(ctx context.Context) ([]entity.Transaction, error) {
	return c.transaction.List(ctx)
}

func (c *transaction) Get(ctx context.Context, filter entity.Transaction) (entity.Transaction, error) {
	return c.transaction.Get(ctx, filter)
}

func (c *transaction) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	return c.transaction.Create(ctx, transaction)
}

func (c *transaction) Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	return c.transaction.Update(ctx, transaction)
}

func (c *transaction) Delete(ctx context.Context, transactionId string) error {
	return c.transaction.Delete(ctx, transactionId)
}
