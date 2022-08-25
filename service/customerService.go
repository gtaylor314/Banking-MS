package service

import (
	"Banking/domain"
	"Banking/dto"
	"Banking/errs"
)

// CustomerService is another "port" (interface)  - users and external sources will interact with the business logic
// via this "port"
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// DefaultCustomerService "connects" to the CustomerRepository interface (in other words, it is dependent on it)
// it bridges the primary port (CustomerService) to the secondary port (CustomerRepository)
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	// if a status query parameter wasn't included, we retrieve all customers
	custStatus := ""
	// if a status query parameter was provided and it was active, we only retrieve active customers (status is 1 in db)
	// if it was inactive, we only retrieve inactive customers (status is 0 in db)
	if status == "active" {
		custStatus = "1"
	}
	if status == "inactive" {
		custStatus = "0"
	}

	cust, err := d.repo.FindAll(custStatus)
	if err != nil {
		return nil, err
	}
	// create a slice of type CustomerResponse with size 0 to begin with
	custResponse := make([]dto.CustomerResponse, 0)

	for _, customer := range cust {
		custResponse = append(custResponse, customer.ToDto())
	}
	return custResponse, nil
}

func (d DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	cust, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	custResponse := cust.ToDto()

	return &custResponse, nil
}

// NewCustomerService instantiates a DefaultCustomerService with a repository
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
