package response

import (
	"time"
	"vaksin-id-be/model"
)

type UserProfile struct {
	NIK          string
	Email        string
	Fullname     string
	PhoneNum     string
	Gender       string
	VaccineCount int
	Age          int
	Address      model.Addresses
}

type AgeUser struct {
	BirthDate time.Time
	Age       int
}

type UserNearbyHealth struct {
	User             UserProfile
	HealthFacilities []HealthResponse
}
