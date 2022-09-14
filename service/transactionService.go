package service

import (
	"time"

	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/domain"
	"github.com/gtaylor314/Banking-MS/dto"
)

// TransactionService is a "port" implemented by the domain
type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

// DefaultTransactionService is an "adapter" that implements the TransactionService "port"
type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

// NewTransaction() takes a NewTransactionRequest and returns a NewTransactionResponse and an error if any
func (d DefaultTransactionService) NewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	// create transaction
	tran := domain.Transaction{
		// TransactionID is auto-generated when the insert command is run on the db
		TransactionID:   "",
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		// TransactionDate is formatted based on the provided layout which matches the layout the db expects
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		AcctAmount:      0,
	}

	// get the current account amount (account balance)
	acctAmount, err := d.repo.GetAcctAmount(tran.AccountID)
	if err != nil {
		return nil, err
	}
	// set AcctAmount in transaction tran to the current account amount (acctAmount)
	tran.AcctAmount = *acctAmount

	// validate the incoming request
	err = req.Validate(*acctAmount)
	if err != nil {
		return nil, err
	}

	// SaveTrans() inserts the transaction into the transactions table and returns a Transaction object with the Transaction ID
	// which auto-generated upon insert
	transaction, err := d.repo.SaveTrans(tran)
	if err != nil {
		return nil, err
	}

	// ToNewTransactionResponseDto() returns a NewTransactionResponse object
	resp := transaction.ToNewTransactionResponseDto()

	return &resp, err
}

// NewTransactionService() takes a TransactionRepository and returns a DefaultTransactionService with the repo property
// initialized to the passed in TransactionRepository
func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repo}
}
