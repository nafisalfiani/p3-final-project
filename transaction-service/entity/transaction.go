package entity

import "time"

type Transaction struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	TicketId    string    `json:"ticket_id"`
	TicketCount int64     `json:"ticket_count"`
	BuyerId     string    `json:"buyer_id"`
	SellerId    string    `json:"seller_id"`
	Amount      float32   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
