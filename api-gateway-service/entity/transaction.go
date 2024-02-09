package entity

import (
	"strings"
	"time"
)

type Transaction struct {
	Id               string    `json:"id"`
	TicketId         string    `json:"ticket_id"`
	TicketCount      int64     `json:"ticket_count"`
	BuyerId          string    `json:"buyer_id"`
	SellerId         string    `json:"seller_id"`
	Amount           float32   `json:"amount"`
	XenditPaymentId  string    `json:"xendit_payment_id"`
	XenditPaymentUrl string    `json:"xendit_payment_url"`
	PaymentMethod    string    `json:"payment_method"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `json:"updated_by"`
}

type TransactionCreateRequest struct {
	TicketId string `json:"ticket_id"`
}

type TransactionUpdateRequest struct {
	Id string `json:"id"`
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

func (x *XenditWebhookBody) GetPaymentId() (paymentType string, paymentId string, err error) {
	// xendit eternal id format is -> ketson-service:<payment_type>:<payment_id>
	res := strings.Split(x.ExternalID, ":")
	paymentType = res[1]
	paymentId = res[2]
	return
}
