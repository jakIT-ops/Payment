package services

import (
	"log"
	"payment_full/db"
	"payment_full/models"
	"payment_full/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

var A = new(Account)

func createRandomAccount(t *testing.T) models.Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	db.ConnectDb()
	log.Println(arg)
	account, err := A.CreateAccount(arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.Id)

	return account
	// models.Account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := A.GetAccount(account1.AccountId)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.AccountId, account2.AccountId)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		Id:      account1.AccountId,
		Balance: utils.RandomMoney(),
	}

	account2, err := A.UpdateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.AccountId, account2.AccountId)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

// func TestDeleteAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	err := A.DeleteAccount(account1.AccountId)
// 	require.NoError(t, err)

// 	account2, err := A.GetAccount(account1.AccountId)
// 	require.Error(t, err)
// 	require.EqualError(t, err, err.Error())
// 	require.Empty(t, account2)
// }
