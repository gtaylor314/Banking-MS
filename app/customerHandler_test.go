package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/dto"
	"github.com/gtaylor314/Banking-MS/mocks/service"
)

var router *mux.Router
var custHandler CustomerHandlers
var mockCustServ *service.MockCustomerService

func setup(t *testing.T) func() {
	// gomock uses a controller which manages the state of the mock - create controller by passing in t
	ctrl := gomock.NewController(t)
	// create a mock customer service object
	mockCustServ = service.NewMockCustomerService(ctrl)
	// define customer handler
	custHandler = CustomerHandlers{
		service: mockCustServ,
	}
	// create a new router using gorilla mux NewRouter()
	router = mux.NewRouter()
	// HandleFunc() registers a new route with a matcher for the URL path
	// since the custHandler's service property holds a MockCustomerService object - when getAllCustomers is called, the
	// mock version of GetAllCustomers will be called within the method
	router.HandleFunc("/customers", custHandler.getAllCustomers)

	// reset the router to allow setup() to be used with multiple tests
	return func() {
		router = nil
	}
}

// TestGetAllCustomersSuccess() should pass when a list of customers and a http.StatusOK code are returned
func TestGetAllCustomersSuccess(t *testing.T) {
	// using AAA format for testing
	// Arrange - setup test
	resetRouter := setup(t)
	// we defer the running of resetRouter() until after test is complete
	defer resetRouter()

	// create a set of dummy customers to test with
	dummyCustList := []dto.CustomerResponse{
		{
			ID:          "1",
			Name:        "Darrow",
			City:        "Lykos",
			Zipcode:     "11111",
			DateofBirth: "12/1/2177",
			Status:      "1",
		},
		{
			ID:          "2",
			Name:        "Mustang",
			City:        "Agea",
			Zipcode:     "22222",
			DateofBirth: "04/15/2178",
			Status:      "1",
		},
	}

	// define expectations for the mock customer service
	mockCustServ.EXPECT().GetAllCustomers("").Return(dummyCustList, nil)

	// create http request
	// we use MethodGet as getAllCustomers is a Get operation, we provide the url path, and no request body
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act - execute test

	// create a response writer using httptest package - ResponseRecorder is an implementation of the response writer
	// recorder will record all responses rather than writing back to the client
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert - test expectations
	// a successful return of all customers should result in a 200 status code (http.StatusOK)
	if recorder.Code != http.StatusOK {
		t.Error("fail while testing the status code 200")
	}
}

// TestGetAllCustomersInternalErr() should pass when no customers and a http.StatusInternalServerError (code 500) are returned
func TestGetAllCustomersInternalErr(t *testing.T) {
	// using AAA format for testing
	// Arrange - setup test
	resetRouter := setup(t)
	// we defer the running of resetRouter() until after test is complete
	defer resetRouter()

	// define expectations for the mock customer service
	mockCustServ.EXPECT().GetAllCustomers("").Return(nil, errs.UnexpectedErr("database error"))

	// create http request
	// we use MethodGet as getAllCustomers is a Get operation, we provide the url path, and no request body
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act - execute test

	// create a response writer using httptest package - ResponseRecorder is an implementation of the response writer
	// recorder will record all responses rather than writing back to the client
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert - test expectations
	// a successful return of all customers should result in a 200 status code (http.StatusOK)
	if recorder.Code != http.StatusInternalServerError {
		t.Error("fail while testing the status code 500")
	}
}
