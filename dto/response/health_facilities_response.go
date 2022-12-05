package response

import (
	"vaksin-id-be/model"
)

type HealthResponse struct {
	ID       string
	Email    string
	PhoneNum string
	Name     string
	Image    *string
	Ranges   int
	Address  model.Addresses
}
