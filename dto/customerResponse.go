package dto

// CustomerResponse will serve as our data transfer object (dto) - this will keep our underlying domain implementation private
// and allow us to omit information we do not want exposed to the public (e.g. removing the date of birth property in the dto)
type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"` // if customer is active or not
}
