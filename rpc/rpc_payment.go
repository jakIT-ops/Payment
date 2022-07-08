package rpc

import (
	"context"
	"errors"
	"log"
	"payment_full/db"
	"payment_full/models"
	"payment_full/pb"
)

// Payment server
type PaymentServer struct {
	pb.UnimplementedTransactionServer
}

func (p *PaymentServer) DebitAmount(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Received: %v", in)
	// var trans models.Transaction
	trans := models.Transaction{AccountId: in.AccountId, TransactionValue: in.TransactionInfor, DebitAmount: in.DebitAmount, CreditAmount: in.CreditAmount}
	db.Database.Db.Create(&trans)
	res := pb.TransactionResponse{TransactionInfor: trans.TransactionValue, TransactionId: int32(trans.Id)}
	return &res, nil
}

func (p *PaymentServer) CreditAmount(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Received: %v", in)
	// var trans models.Transaction
	trans := models.Transaction{AccountId: in.AccountId, TransactionValue: in.TransactionInfor, DebitAmount: in.DebitAmount, CreditAmount: in.CreditAmount}
	db.Database.Db.Create(&trans)
	res := pb.TransactionResponse{TransactionInfor: trans.TransactionValue, TransactionId: int32(trans.Id)}
	return &res, nil
}

// find account
func findAccount(id string, account *models.Account) error {
	db.Database.Db.Find(&account, "account_id = ?", id)

	if account.AccountId == "" {
		return errors.New("account does not exist")
	}
	return nil
}

func (p *PaymentServer) GetBalance(ctx context.Context, in *pb.Account_Id) (*pb.Balance, error) {
	log.Printf("Received: %v", in)
	var res pb.Balance
	var creditAmount int64
	var debitAmount int64
	var account models.Account

	db.Database.Db.Find(&account, "account_id = ?", in.Value)

	db.Database.Db.Raw("SELECT SUM(debit_amount) FROM transactions WHERE account_id = ?", in.Value).Scan(&debitAmount)
	db.Database.Db.Raw("SELECT SUM(credit_amount) FROM transactions WHERE account_id = ?", in.Value).Scan(&creditAmount)

	res.Amount = creditAmount - debitAmount

	findAccount(in.Value, &account)

	account.Balance = res.Amount

	db.Database.Db.Save(&account)

	return &res, nil
}
