package xendit

import (
	"context"
	"fmt"

	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/transaction-service/entity"
	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/invoice"
)

type xndClient struct {
	logger log.Interface
	xnd    *xendit.APIClient
}

type Interface interface {
	CreatePayment(ctx context.Context, payment entity.XenditPaymentRequest) (entity.XenditPaymentResponse, error)
	CheckPayment(ctx context.Context, payment entity.XenditPaymentResponse) (entity.XenditPaymentResponse, error)
}

// Init create xendit repository
func Init(xnd *xendit.APIClient, logger log.Interface) Interface {
	return &xndClient{
		logger: logger,
		xnd:    xnd,
	}
}

func (x *xndClient) CreatePayment(ctx context.Context, payment entity.XenditPaymentRequest) (entity.XenditPaymentResponse, error) {
	req := x.xnd.InvoiceApi.CreateInvoice(ctx).CreateInvoiceRequest(invoice.CreateInvoiceRequest{
		ExternalId:      fmt.Sprintf("ketson-service%v", payment.PaymentId),
		Amount:          payment.Amount,
		Description:     payment.InvoiceDescription,
		InvoiceDuration: payment.InvoiceExpiry,
		Customer: &invoice.CustomerObject{
			GivenNames: *invoice.NewNullableString(payment.InvoiceName),
			Email:      *invoice.NewNullableString(payment.InvoiceEmail),
		},
		Currency: payment.Currency,
		Items:    x.invoiceItems(payment.Items),
	})

	xenditResp, res, err := req.Execute()
	if err != nil {
		x.logger.Error(ctx, res)
		x.logger.Error(ctx, fmt.Sprintf("Error when calling `PaymentRequestApi.CreatePaymentRequest``: %#v\n", err))
		return entity.XenditPaymentResponse{}, err
	}

	resp := entity.XenditPaymentResponse{
		XenditPaymentId:   *xenditResp.Id,
		PaymentId:         xenditResp.ExternalId,
		InvoiceExpiryDate: xenditResp.ExpiryDate,
		InvoiceStatus:     xenditResp.Status.String(),
		InvoiceAmount:     xenditResp.Amount,
		InvoiceUrl:        xenditResp.InvoiceUrl,
	}

	if xenditResp.PaymentMethod != nil && xenditResp.PaymentMethod.IsValid() {
		resp.PaymentMethod = xenditResp.PaymentMethod.String()
	}

	return resp, nil
}

func (x *xndClient) invoiceItems(items []entity.PaymentItems) []invoice.InvoiceItem {
	invoiceItems := []invoice.InvoiceItem{}
	for i := range items {
		invoiceItems = append(invoiceItems, invoice.InvoiceItem{
			Name:     items[i].Name,
			Price:    items[i].Price,
			Quantity: items[i].Quantity,
		})
	}

	return invoiceItems
}

func (x *xndClient) CheckPayment(ctx context.Context, payment entity.XenditPaymentResponse) (entity.XenditPaymentResponse, error) {
	req := x.xnd.InvoiceApi.GetInvoiceById(ctx, payment.XenditPaymentId)
	xenditResp, res, err := req.Execute()
	if err != nil {
		x.logger.Error(ctx, res)
		x.logger.Error(ctx, fmt.Sprintf("Error when calling `PaymentRequestApi.CheckPayment``: %#v\n", err))
		return entity.XenditPaymentResponse{}, err
	}

	resp := entity.XenditPaymentResponse{
		XenditPaymentId:   *xenditResp.Id,
		PaymentId:         xenditResp.ExternalId,
		InvoiceExpiryDate: xenditResp.ExpiryDate,
		InvoiceStatus:     xenditResp.Status.String(),
		InvoiceAmount:     xenditResp.Amount,
		InvoiceUrl:        xenditResp.InvoiceUrl,
	}

	if xenditResp.PaymentMethod != nil && xenditResp.PaymentMethod.IsValid() {
		resp.PaymentMethod = xenditResp.PaymentMethod.String()
	}

	return resp, nil
}
