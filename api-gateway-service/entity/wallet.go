package entity

import "time"

type Wallet struct {
	Id        string          `json:"id"`
	UserId    string          `json:"user_id"`
	Balance   float32         `json:"balance"`
	History   []WalletHistory `json:"history"`
	CreatedAt time.Time       `json:"created_at"`
	CreatedBy string          `json:"created_by"`
	UpdatedAt time.Time       `json:"updated_at"`
	UpdatedBy string          `json:"updated_by"`
}

type WalletHistory struct {
	Id              string    `json:"id"`
	WalletId        string    `json:"wallet_id"`
	PreviousBalance float32   `json:"previous_balance"`
	CurrentBalance  float32   `json:"current_balance"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
}
