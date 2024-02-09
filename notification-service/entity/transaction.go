package entity

import "time"

type Transaction struct {
	Id               string    `json:"id" gorm:"primaryKey"`
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
