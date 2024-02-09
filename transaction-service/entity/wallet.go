package entity

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id        uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid()"`
	UserId    string          `json:"user_id" gorm:"not null"`
	Balance   float32         `json:"balance" gorm:"not null"`
	History   []WalletHistory `json:"history" gorm:"not null"`
	CreatedAt time.Time       `json:"created_at" bson:"created_at,omitempty"`
	CreatedBy string          `json:"created_by" bson:"created_by,omitempty"`
	UpdatedAt time.Time       `json:"updated_at" bson:"updated_at,omitempty"`
	UpdatedBy string          `json:"updated_by" bson:"updated_by,omitempty"`
}

type WalletHistory struct {
	Id              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	WalletId        string    `json:"wallet_id" gorm:"not null"`
	PreviousBalance float32   `json:"previous_balance" gorm:"not null"`
	CurrentBalance  float32   `json:"current_balance" gorm:"not null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at,omitempty"`
	CreatedBy       string    `json:"created_by" bson:"created_by,omitempty"`
}
