package rpc

import (
	"context"
	"payment_full/db"
	"payment_full/pb"
	"payment_full/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

var P = new(PaymentServer)

// pb.TransactionRequest
func TestCreateNewRandomTransaction(t *testing.T) {
	arg := pb.TransactionRequest{
		AccountId:        utils.RandomOwner(),
		TransactionInfor: "credit&debit",
		CreditAmount:     utils.RandomMoney(),
		DebitAmount:      utils.RandomMoney(),
	}
	db.ConnectDb()
	transaction, err := P.CreditAmount(context.Background(), &arg)

	require.NoError(t, err)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	// require.Equal(t, arg.AccountId, transaction.TransactionId)
	require.Equal(t, arg.TransactionInfor, transaction.TransactionInfor)

	require.NotZero(t, transaction.TransactionId)
	// return arg
}

// func TestCreateTransaction(t *testing.T) {
// 	createNewRandomTransaction(t)
// }
