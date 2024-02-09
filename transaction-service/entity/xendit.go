package entity

import (
	"strconv"
	"strings"
	"time"
)

var (
	IdrCurrency = "IDR"

	DescriptionTicketTransaction = "Rental Payment"

	InvoiceExpiry = "86400"
)

const (
	InvoiceStatusPending = "PENDING"
	InvoiceStatusPaid    = "PAID"

	PaymentMethodWaiting = "WAITING"

	PaymentTypeTicketTransaction = "ticket-transaction"
)

type XenditPaymentRequest struct {
	PaymentId          string
	Amount             float64
	PaymentMethod      string
	Currency           *string
	InvoiceDescription *string
	InvoiceExpiry      *string
	InvoiceName        *string
	InvoiceEmail       *string
	Items              []PaymentItems
}

type PaymentItems struct {
	Name     string
	Price    float32
	Quantity float32
}

type XenditPaymentResponse struct {
	XenditPaymentId   string    `json:"xendit_payment_id"`
	PaymentId         string    `json:"payment_id"`
	InvoiceExpiryDate time.Time `json:"expiry_date"`
	InvoiceStatus     string    `json:"status"`
	InvoiceAmount     float64   `json:"amount"`
	InvoiceUrl        string    `json:"url"`
	PaymentMethod     string    `json:"payment_method"`
}

func (x *XenditPaymentResponse) GetPaymentId() (paymentType string, paymentId int, err error) {
	res := strings.Split(x.PaymentId, ":")
	paymentType = res[1]
	paymentId, err = strconv.Atoi(res[2])
	return
}

type XenditCheckPayment struct {
	XenditPaymentId string `json:"xendit_payment_id"`
	PaymentId       string `json:"payment_id"`
}

type XenditWebhookBody struct {
	ID                     string    `json:"id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 int       `json:"amount"`
	PaidAmount             int       `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 time.Time `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	AdjustedReceivedAmount int       `json:"adjusted_received_amount"`
	FeesPaidAmount         int       `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
}

func (x *XenditWebhookBody) GetPaymentId() (paymentType string, paymentId int, err error) {
	// xendit eternal id format is -> ketson-service:<payment_type>:<payment_id>
	res := strings.Split(x.ExternalID, ":")
	paymentType = res[1]
	paymentId, err = strconv.Atoi(res[2])
	return
}
