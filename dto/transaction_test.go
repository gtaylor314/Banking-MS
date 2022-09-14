package dto

import (
	"net/http"
	"testing"
)

// TestTransactionType() - error when the transaction type is not deposit or withdrawal
func TestTransactionType(t *testing.T) {
	// following the AAA format for testing
	// Arrange - setup test
	tranRequest := NewTransactionRequest{
		// the transaction type should be deposit or withdrawal so we will use "invalid" to cause an error
		TransactionType: "invalid",
	}

	// Act - execute test
	// here I pass in a dumby account amount (aka account balance)
	appError := tranRequest.Validate(1000)

	// Assert - test expectations
	// we expect the error message below from the Validate method - if another error message is returned, the test should fail
	if appError.Message != "error: transaction type must be either withdrawal or deposit" {
		t.Error("invalid message while testing the transaction type")
	}

	// we expect code 422 UnprocessableEntity from the Validate method - if another code is returned, the test should fail
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("invalid code while testing the transaction type")
	}
}

// TestTransactionAmount() - error when the transaction amount is less than zero
func TestTransactionAmount(t *testing.T) {
	// following the AAA format for testing
	// Arrange - setup test
	tranRequest := NewTransactionRequest{
		// the amount should be greater than or equal to zero - we use negative one to elicit error
		Amount: -1,
		// need a valid transaction type for test to complete
		TransactionType: "deposit",
	}
	// Act - execute test
	// here I pass in a dumby account amount (aka account balance)
	appError := tranRequest.Validate(1000)

	// Assert - text expectations
	// we expect the below error message from Validate() - if another error message is returned, the test should fail
	if appError.Message != "error: transaction amount must be zero or greater" {
		t.Error("invalid message while testing the transaction amount")
	}

	// we expect code 422 UnprocessableEntity from Validate() - if another code is returned, the test should fail
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("invalid code while testing the transaction amount")
	}

}

// TestTransactionAmountValue() - error if the transaction type is withdrawal AND the transaction amount is greater than
// the account amount (aka account balance)
func TestTransactionAmountValue(t *testing.T) {
	// following the AAA format for testing
	// Arrange - setup testing
	tranRequest := NewTransactionRequest{
		Amount:          10000,
		TransactionType: "withdrawal",
	}

	// Act - execute test
	// here I pass in an account amount (aka account balance) less than the transaction amount
	appError := tranRequest.Validate(1000)

	// Assert - test expectations
	// we expect the below error message from Validate() - if a different message is returned, the test should fail
	if appError.Message != "error: transaction amount must be less than or equal to account's balance" {
		t.Error("invalid message while testing the transaction amount against the account amount")
	}

	// we expect code 422 UnprocessableEntity from Validate() - if a different code is returned, the test should fail
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("invalid code while testing the transaction amount against the account amount")
	}
}
