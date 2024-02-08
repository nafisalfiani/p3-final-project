package entity

import (
	"time"
)

type User struct {
	Id              string    `json:"id,omitempty"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"-"`
	Password        string    `json:"password"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       string    `json:"updated_by"`
}
