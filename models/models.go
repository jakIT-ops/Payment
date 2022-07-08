package models

import (
	"time"
)

type Account struct {
	Id        int32     `gorm:"primaryKey"`
	AccountId string    `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Transaction struct {
	Id               uint      `gorm:"primaryKey"`
	AccountId        string    `json:"id"`
	TransactionValue string    `json:"value"`
	DebitAmount      int64     `json:"debit_amount"`
	CreditAmount     int64     `json:"credit_amount"`
	CreatedAt        time.Time `json:"created_at"`
}

type User struct {
	Id       uint   `gorm:"primaryKey"`
	UserName string `json:"username"`
	// Role              string    `json:"role"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
