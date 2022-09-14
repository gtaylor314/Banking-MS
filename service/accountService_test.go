package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	// we use an alias for the actual domain package since our mock also has a domain package
	"github.com/gtaylor314/Banking-Lib/errs"
	realdomain "github.com/gtaylor314/Banking-MS/domain"
	"github.com/gtaylor314/Banking-MS/dto"
	"github.com/gtaylor314/Banking-MS/mocks/domain"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

// setup() sets up the mock controller, mock repository, account service and resets the service to nil
func setup(t *testing.T) func() {
	// create a controller for the mock repository - the controller manages the state of the mock
	ctrl := gomock.NewController(t)
	// create a mock repository
	mockRepo = domain.NewMockAccountRepository(ctrl)
	// create a service
	service = NewAccountService(mockRepo)
	// returned function resets the service back to nil
	return func() {
		service = nil
	}
}

// TestNewAccountValidationErr() should pass when the new account request fails validation and returns an error
func TestNewAccountValidationErr(t *testing.T) {
	// using AAA format for testing
	// Arrange - setup test
	// create an account request
	acctReq := dto.NewAccountRequest{
		CustomerID: "1",
		// must be saving or checking - can set to "invalid" to cause the validation to fail
		AccountType: "saving",
		// must deposit 5,000.00 or more to create a new account - 1000.00 will cause the validation to fail
		Amount: 1000.00,
	}

	// create a service
	// NewAccountService() generally takes an account repository; however, since the validation should fail and the test
	// pass before the repository is used within the NewAccount() method, we simply pass nil instead of creating a repo which
	// isn't needed
	service := NewAccountService(nil)

	// Act - execute test
	_, appError := service.NewAccount(acctReq)

	// Assert - test expectations
	// we expect an error from the new account request validation, if it is nil, the test has failed
	if appError == nil {
		t.Error("failed while testing new account request validation")
	}
}

// TestNewAccountFailedSave() should pass when a server side error is returned after calling Save() - meaning new account
// cannot be created
func TestNewAccountFailedSave(t *testing.T) {
	// using AAA format for testing
	// Arrange - setup test
	resetService := setup(t)
	defer resetService()
	// create a new account request with valid details
	acctReq := dto.NewAccountRequest{
		CustomerID:  "1",
		AccountType: "saving",
		Amount:      5000.00,
	}
	// create a test account using account request details
	account := realdomain.Account{
		AccountID:  "", // AccountID is populated once the account is created
		CustomerID: acctReq.CustomerID,
		// OpeningDate is formatted based on the provided layout which matches the layout the db expects
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: acctReq.AccountType,
		Amount:      acctReq.Amount,
		// Status is active (or 1) by default when opening a new account
		Status: "1",
	}
	// define expectations
	mockRepo.EXPECT().Save(account).Return(nil, errs.UnexpectedErr("Unexpected database error"))

	// Act - execute test
	_, appErr := service.NewAccount(acctReq)

	// Assert - test expectations
	// we expect an error, in which, if appErr is nil, then the test has failed
	if appErr == nil {
		t.Error("failed while testing error during new account creation (Save)")
	}
}

// TestNewAccountSuccess() should pass when no errors are returned and the account has been created
func TestNewAccountSuccess(t *testing.T) {
	// using AAA format for testing
	// Arrange - setup test
	resetService := setup(t)
	defer resetService()

	// create a new account request with valid details
	acctReq := dto.NewAccountRequest{
		CustomerID:  "1",
		AccountType: "saving",
		Amount:      5000.00,
	}
	// create a test account using account request details
	account := realdomain.Account{
		AccountID:  "", // AccountID is populated once the account is created
		CustomerID: acctReq.CustomerID,
		// OpeningDate is formatted based on the provided layout which matches the layout the db expects
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: acctReq.AccountType,
		Amount:      acctReq.Amount,
		// Status is active (or 1) by default when opening a new account
		Status: "1",
	}
	// create a test account with an account id to simulate a created account
	acctWithID := account
	acctWithID.AccountID = "101"
	// define expectations
	mockRepo.EXPECT().Save(account).Return(&acctWithID, nil)

	// Act - execute test
	returnedAcct, appErr := service.NewAccount(acctReq)

	// Assert - test expectations
	// we expect no errors during a successful account creation
	if appErr != nil {
		t.Error("failed while creating new account - received error")
	}
	// we expect the returned account to have an account ID which matches the test account's account ID
	if returnedAcct.AccountID != acctWithID.AccountID {
		t.Error("failed while creating new account - account ID does not match")
	}
}
