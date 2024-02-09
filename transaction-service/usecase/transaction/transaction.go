package transaction

import (
	"context"
	"fmt"

	"github.com/nafisalfiani/p3-final-project/lib/broker"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	transactionDom "github.com/nafisalfiani/p3-final-project/transaction-service/domain/transaction"
	"github.com/nafisalfiani/p3-final-project/transaction-service/domain/xendit"
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
	logger      log.Interface
	transaction transactionDom.Interface
	xendit      xendit.Interface
	broker      broker.Interface
}

func Init(logger log.Interface, prd transactionDom.Interface, xnd xendit.Interface, broker broker.Interface) Interface {
	return &transaction{
		logger:      logger,
		transaction: prd,
		xendit:      xnd,
		broker:      broker,
	}
}

func (c *transaction) List(ctx context.Context) ([]entity.Transaction, error) {
	return c.transaction.List(ctx)
}

func (c *transaction) Get(ctx context.Context, filter entity.Transaction) (entity.Transaction, error) {
	return c.transaction.Get(ctx, filter)
}

func (c *transaction) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	trx, err := c.transaction.Create(ctx, transaction)
	if err != nil {
		return trx, err
	}

	paymentType := entity.PaymentTypeTicketTransaction
	invoiceName := "customer"
	invoiceEmail := "nafisa.alfiani.ica@gmail.com"
	resp, err := c.xendit.CreatePayment(ctx, entity.XenditPaymentRequest{
		PaymentId:          fmt.Sprintf(":ticket-transaction:%v", trx.Id),
		Amount:             float64(trx.Amount),
		InvoiceName:        &invoiceName,
		InvoiceEmail:       &invoiceEmail,
		InvoiceDescription: &paymentType,
		InvoiceExpiry:      &entity.InvoiceExpiry,
		Currency:           &entity.IdrCurrency,
	})
	if err != nil {
		return entity.Transaction{}, err
	}

	trx.XenditPaymentId = resp.XenditPaymentId
	trx.XenditPaymentUrl = resp.InvoiceUrl
	trx.PaymentMethod = entity.PaymentMethodWaiting
	trx.Status = entity.InvoiceStatusPending
	trx.Type = entity.PaymentTypeTicketTransaction

	newTrx, err := c.transaction.Update(ctx, trx)
	if err != nil {
		return newTrx, err
	}

	if err := c.broker.PublishMessage(entity.TopicNewTransaction, newTrx); err != nil {
		c.logger.Error(ctx, err)
	}

	return newTrx, nil
}

func (c *transaction) Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	return c.transaction.Update(ctx, transaction)
}

func (c *transaction) Delete(ctx context.Context, transactionId string) error {
	return c.transaction.Delete(ctx, transactionId)
}
