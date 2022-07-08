package services

import (
	"payment_full/models"
)

type AccountServices interface {
	// Account
	AddAccountBalance(arg GetAccountBalanceParams) int64
	CreateAccount(arg CreateAccountParams) (models.Account, error)
	DeleteAccount(id string) error
	GetAccount(id string) (models.Account, error)
	ListAccounts() ([]models.Account, error)
	UpdateAccount(arg UpdateAccountParams) (models.Account, error)
	// User
	CreateUser(arg CreateUserParams) (models.User, error)
	GetUser(username string) (models.User, error)
}
