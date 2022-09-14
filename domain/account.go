package domain

import (
	"time"

	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/dto"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

// AccountRepository is a "port" implemented by the server side - AccountRepositoryDb is an adapater which implements the
// AccountRepository interface
//
//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/gtaylor314/Banking-MS/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (acct Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountID: acct.AccountID}
}

// NewAccountNoID() returns an account object without an account ID - the account ID is auto-generated when the account is
// created/saved to the database - Save() populates the account with the auto-generated account ID
func NewAccountNoID(customerID string, accountType string, amount float64) Account {
	return Account{
		AccountID:   "",
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1", // new accounts are active (status = 1) by default
	}
}
