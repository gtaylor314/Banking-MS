package domain

import (
	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/dto"
)

type Transaction struct {
	TransactionID   string
	AccountID       string
	Amount          float64
	TransactionType string
	TransactionDate string
	AcctAmount      float64 // AcctAmount is the account's balance
}

// TransactionRepository is a "port" - implemented by the server side "adapter" TransactionRepositoryDb
type TransactionRepository interface {
	SaveTrans(Transaction) (*Transaction, *errs.AppError)
	GetAcctAmount(string) (*float64, *errs.AppError)
}

// ToNewTransactionResponseDto() takes a transaction object and returns a NewTransactionResponse object
func (tran Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionID: tran.TransactionID,
		AcctAmount:    tran.AcctAmount,
	}
}
