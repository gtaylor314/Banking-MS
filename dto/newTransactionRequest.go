package dto

import (
	"strings"

	"github.com/gtaylor314/Banking-Lib/errs"
)

// NewTransactionRequest is a dto which provides customer sourced data to the domain for transaction creation
type NewTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (req NewTransactionRequest) Validate(acctAmount float64) *errs.AppError {
	// confirm the transaction type is either withdrawal or deposit
	// strings.ToLower() ensures that the transaction type is all lower case prior to comparison
	if strings.ToLower(req.TransactionType) != "withdrawal" && strings.ToLower(req.TransactionType) != "deposit" {
		return errs.ValidationErr("error: transaction type must be either withdrawal or deposit")
	}

	// confirm the transaction amount is positive (0 or greater)
	if req.Amount < 0 {
		return errs.ValidationErr("error: transaction amount must be zero or greater")
	}

	// confirm, if transaction type is withdrawl, that the account amount is greater than or equal to the amount of the
	// transaction
	if strings.ToLower(req.TransactionType) == "withdrawal" {
		if req.Amount > acctAmount {
			return errs.ValidationErr("error: transaction amount must be less than or equal to account's balance")
		}
	}

	return nil
}
