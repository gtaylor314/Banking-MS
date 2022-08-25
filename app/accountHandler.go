package app

import (
	"Banking/dto"
	"Banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (acctHandler AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	// the customer_id is provided as a query parameter (part of the URL)
	// Vars() returns the route variables, if any, in the current request
	vars := mux.Vars(r)
	var req dto.NewAccountRequest
	// we populate the customer_id in our NewAccountRequest
	req.CustomerID = vars["customer_id"]
	// the account_type and amount for the new account are provided via the JSON body - here we create a new decoder which
	// reads the response body and stores the data in the NewAccountRequest (req)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// if there is an error, it must be a bad request
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, appErr := acctHandler.service.NewAccount(req)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.MessageOnly())
		return
	}
	writeResponse(w, http.StatusCreated, resp)
}
