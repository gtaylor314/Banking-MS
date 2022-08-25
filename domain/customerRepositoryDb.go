package domain

import (
	"database/sql"

	"github.com/gtaylor314/Banking-MS/errs"
	"github.com/gtaylor314/Banking-MS/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CustomerRepositoryDb represents an "adapter" which will "connect" to our "port" which is CustomerRepository
type CustomerRepositoryDb struct {
	db_conn *sqlx.DB
}

// FindAll() implementation for type CustomerRepositoryDb - returns a slice of Customer objects and an error
func (c CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// create a slice of customers with size zero to begin with
	customers := make([]Customer, 0)
	// create empty query string
	var findCustSql string
	// create an error variable
	var err error

	// define string for MySQL query - this is used if the status query parameter was left blank (we retrieve all customers)
	if status == "" {
		findCustSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// Select() both queries the database using the MySQL query (findCustSql) and scans the results into the
		// destination (customers)
		err = c.db_conn.Select(&customers, findCustSql)
	}
	// if status is equal to 1 or to 0 (active or inactive), we change the MySQL query to include the status
	if status == "1" || status == "0" {
		findCustSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		// Select() both queries the database using the MySQL query (findCustSql with status) and it scans the results into the
		// destination (customers)
		err = c.db_conn.Select(&customers, findCustSql, status)
	}

	if err != nil {
		// err.Error() takes the error and returns it as a string
		logger.Error("error during query of customers table " + err.Error())
		return nil, errs.UnexpectedErr("unexpected database error")
	}

	return customers, nil
}

// FindById() implementation for type CustomerRepositoryDb - returns a pointer to a Customer object so we can
// use nil if the id is not found
func (c CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	// define string for MySQL query
	customerIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var cust Customer
	// Get() both queries the database for the row using the MySQL query (customerIdSql) and scans the result into
	// the destination
	err := c.db_conn.Get(&cust, customerIdSql, id)
	if err != nil {
		// err may be due to a request for a customer that doesn't exist or an issue with the database/scan method
		// if err == sql.ErrNoRows, then the customer doesn't exist
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundErr("customer not found")
		}
		// otherwise, the error is internal - either with the database
		logger.Error("error during query of customer table " + err.Error())
		// custom error for a better user's experience
		return nil, errs.UnexpectedErr("unexpected database error")
	}
	return &cust, nil
}

// NewCustomerRepositoryDb() is a helper function that creates a connection to the database
func NewCustomerRepositoryDb(db_conn *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{db_conn: db_conn}
}
