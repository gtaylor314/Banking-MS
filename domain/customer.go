package domain

import (
	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/dto"
)

// Customer represents our business object - a customer
// Since we are using sqlx - we use db tags to match the column names in the database with the property name in the struct
type Customer struct {
	ID          string `db:"customer_id"`
	Name        string // column name matches property name already
	City        string // column name matches property name already
	Zipcode     string // column name matches property name already
	DateofBirth string `db:"date_of_birth"`
	// if customer is active or not
	Status string // column name matches property name already
}

// CustomerRepository represents our "port" (interface) for the server side to interact with our business object
type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	// FindById returns a pointer to a Customer object so that we may return nil if the ID is not found
	FindById(string) (*Customer, *errs.AppError)
}

func (cust Customer) statusText() string {
	// want to convert the internal status (0 or 1) to a user friendly status of inactive or active
	statusText := "active"
	if cust.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

// ToDto converts a domain.Customer object to a dto.CustomerResponse object - this makes the underlying implementation of our
// business domain private
func (cust Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          cust.ID,
		Name:        cust.Name,
		City:        cust.City,
		Zipcode:     cust.Zipcode,
		DateofBirth: cust.DateofBirth,
		Status:      cust.statusText(),
	}
}
