package app

import (
	"encoding/json"
	"net/http"

	"github.com/gtaylor314/Banking-MS/dto"
	"github.com/gtaylor314/Banking-MS/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (tranHandler TransactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {
	// creating a NewTransactionRequest
	var req dto.NewTransactionRequest
	// the account id, amount, and transaction type for the new transaction are provided via the JSON body - here we create
	// a new decoder which reads the response body and stores the data in the NewTransactionRequest (req)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// something with the request is incorrect
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// if request is good, create a new transaction with the req
	resp, appErr := tranHandler.service.NewTransaction(req)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.MessageOnly())
		return
	}
	// write the NewTransactionResponse to the ResponseWriter
	writeResponse(w, http.StatusCreated, resp)
}
