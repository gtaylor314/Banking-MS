package app

import (
	"encoding/json"
	"net/http"

	"github.com/gtaylor314/Banking-MS/service"

	"github.com/gorilla/mux"
)

// CustomerHandler "connects" to the CustomerService interface (in other words, it is dependent on it)
type CustomerHandlers struct {
	service service.CustomerService
}

func (custHandler *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	// check if there is a query parameter (specifically for key "status")
	queryParam := r.URL.Query().Get("status")
	// grab the list of customers and an error if any
	customers, err := custHandler.service.GetAllCustomers(queryParam)
	if err != nil {
		// if there are errors, set response header content-type, status code, and response body via writeResponse()
		// and return
		writeResponse(w, err.Code, err.MessageOnly())
		return
	}
	// if there are no errors, set response header content-type, status code, and response body via writeResponse()
	writeResponse(w, http.StatusOK, customers)
	return

}

func (custHandler *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	// grab the customer_id from the http.Request
	// mux.Vars() is provided by the gorilla mux package - it returns the route variables for the current http.Request, if any
	vars := mux.Vars(r)
	customer, err := custHandler.service.GetCustomer(vars["customer_id"])
	if err != nil {
		// customer_id was not found
		writeResponse(w, err.Code, err.MessageOnly())
		return
	}
	// customer_id was found
	writeResponse(w, http.StatusOK, customer)
}

// writeResponse sets our response header (content-type to application/json), the status code for the response header, and
// finally encodes the response body in the json format (data interface{} allows us to pass anything for the response body)
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
