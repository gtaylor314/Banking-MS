package dto

import (
	"strings"

	"github.com/gtaylor314/Banking-Lib/errs"
)

// NewAccountRequest is a dto which provides customer generated data to the domain for account creation
type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

// Validate() confirms that the NewAccountRequest meets all criteria for creating a new account
func (req NewAccountRequest) Validate() *errs.AppError {
	if req.Amount < 5000.00 {
		return errs.ValidationErr("error: must deposit at least 5,000.00 to create new account")
	}
	// strings.ToLower ensures that req.AccountType is all lower case when compared against "saving" or "checking"
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return errs.ValidationErr("error: account type must be either saving or checking")
	}
	return nil
}
