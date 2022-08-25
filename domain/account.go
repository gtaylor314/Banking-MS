package domain

import (
	"github.com/gtaylor314/Banking-MS/dto"
	"github.com/gtaylor314/Banking-MS/errs"
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
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (acct Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: acct.AccountID}
}
