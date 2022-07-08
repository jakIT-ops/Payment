package services

import (
	"errors"
	"log"
	"payment_full/db"
	"payment_full/models"

	"github.com/google/uuid"
)

type Account struct {
	AccountId string `json:"id"`
	Owner     string `json:"owner"`
	Currency  string `json:"currency"`
}

type GetAccountBalanceParams struct {
	Id string `json:"id"`
}

func (a *Account) GetAccountBalance(arg GetAccountBalanceParams) int64 {
	id := arg.Id
	var amount int64
	var creditAmount int64
	var debitAmount int64

	db.Database.Db.Raw("SELECT SUM(debit_amount) FROM transactions WHERE account_id = ?", id).Scan(&debitAmount)
	db.Database.Db.Raw("SELECT SUM(credit_amount) FROM transactions WHERE account_id = ?", id).Scan(&creditAmount)

	amount = creditAmount - debitAmount
	return amount
}

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (a Account) CreateAccount(arg CreateAccountParams) (models.Account, error) {
	var newAccount models.Account

	newAccount.AccountId = uuid.New().String()
	newAccount.Owner = arg.Owner
	newAccount.Currency = arg.Currency
	newAccount.Balance = int64(arg.Balance)

	db.Database.Db.Create(&newAccount)

	// var res = Account{AccountId: newAccount.AccountId, Owner: newAccount.Owner, Currency: newAccount.Currency}
	return newAccount, nil
}

type UpdateAccountParams struct {
	Id      string `json:"id"`
	Balance int64  `json:"amount"`
}

func (a Account) ListAccounts() ([]models.Account, error) {
	accounts := []models.Account{}
	// res := []Account{}
	db.Database.Db.Find(&accounts)

	// for _, account := range accounts {
	// 	res = append(res, Account{AccountId: account.AccountId, Owner: account.Owner, Currency: account.Currency})
	// }
	return accounts, nil
}

// find account
func findAccount(id string, account *models.Account) error {
	db.Database.Db.Find(&account, "account_id = ?", id)

	if account.AccountId == "" {
		return errors.New("account does not exist")
	}
	return nil
}

func (a Account) GetAccount(id string) (models.Account, error) {
	var account models.Account
	if err := findAccount(id, &account); err != nil {
		return account, errors.New(err.Error())
	}
	// res := Account{AccountId: account.AccountId, Owner: account.Owner, Currency: account.Currency}
	return account, nil
}

func (a Account) UpdateAccount(arg UpdateAccountParams) (models.Account, error) {
	log.Printf("Received: %v", arg)
	var updateAccount models.Account
	err := findAccount(arg.Id, &updateAccount)
	if err != nil {
		return updateAccount, errors.New(err.Error())
	}

	updateAccount.Balance = arg.Balance

	db.Database.Db.Save(&updateAccount)
	// var res = Account{AccountId: updateAccount.AccountId, Owner: updateAccount.Owner, Currency: updateAccount.Currency}

	return updateAccount, nil
}

func (a Account) DeleteAccount(id string) error {
	var deletedAccount models.Account
	err := findAccount(id, &deletedAccount)
	if err != nil {
		return errors.New(err.Error())
	}
	if err := db.Database.Db.Delete(&deletedAccount).Error; err != nil {
		log.Printf("failed to deleted account %v", err.Error())
	}
	return nil
}
