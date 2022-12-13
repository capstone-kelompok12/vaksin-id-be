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
	Ranges   float64
	Address  model.Addresses
	Session  []model.Sessions
}

type UpdateHealthFacilities struct {
	Email          string
	PhoneNum       string
	Name           string
	Image          *string
	CurrentAddress string
	District       string
	City           string
	Province       string
	Latitude       float64
	Longitude      float64
}
