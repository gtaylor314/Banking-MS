package domain

// CustomerRepositoryStub represents an "adapter" that will "connect" to our "port" which is CustomerRepository
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll implementation for type CustomerRepositoryStub
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	// return the list of customers and a nil error
	return c.customers, nil
}

// NewCustomerRepositoryStub is a helper function which creates a new customer repository stub
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	// creating dummy customers
	customers := []Customer{
		{
			ID:          "1001",
			Name:        "Garrett",
			City:        "Bradenton",
			Zipcode:     "34205",
			DateofBirth: "10-01-1989",
			Status:      "1", // one for active and zero for inactive
		},
		{
			ID:          "1002",
			Name:        "Kimberly",
			City:        "Seattle",
			Zipcode:     "98101",
			DateofBirth: "06-01-1992",
			Status:      "1",
		},
		{
			ID:          "1003",
			Name:        "Noah",
			City:        "Orlando",
			Zipcode:     "32789",
			DateofBirth: "03-01-2021",
			Status:      "1",
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
