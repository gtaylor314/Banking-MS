package service

import (
	"time"

	"github.com/gtaylor314/Banking-MS/domain"
	"github.com/gtaylor314/Banking-MS/dto"
	"github.com/gtaylor314/Banking-MS/errs"
)

// AccountService is a "port" implemented by the domain
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// DefaultAccountService is an "adapter" that implements the AccountService "port"
type DefaultAccountService struct {
	repo domain.AccountRepository
}

// NewAccount() takes an account request dto, populates the account object
func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// validate the incoming request
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// create account
	acct := domain.Account{
		// AccountID is auto-generated when the insert command is run on the db
		AccountID:  "",
		CustomerID: req.CustomerID,
		// OpeningDate is formatted based on the provided layout which matches the layout the db expects
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		// Status is active (or 1) by default when opening a new account
		Status: "1",
	}
	// Save() inserts the account into the accounts table and returns an account object with the AccountID which auto-generated
	// upon insert
	account, err := d.repo.Save(acct)
	if err != nil {
		return nil, err
	}

	// the account object needs to be converted to a dto prior to responding to the user
	resp := account.ToNewAccountResponseDto()

	return &resp, nil
}

// NewAccountService() takes an account repository and creates a new DefaultAccountService
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
