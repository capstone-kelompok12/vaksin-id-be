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
	BirthDate    time.Time
	Age          int
	Address      *model.Addresses
}
type UserHistory struct {
	NIK          string
	Email        string
	Fullname     string
	PhoneNum     string
	Gender       string
	VaccineCount int
	BirthDate    time.Time
	Age          int
	Address      *model.Addresses
	History      []HistoryCustomUser
}

type HistoryCustomUser struct {
	ID         string
	IdBooking  string
	NikUser    string
	IdSameBook string
	Status     *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Booking    BookingHistoryLoop
}

type BookingHistoryLoop struct {
	ID               string
	IdSession        string
	NikUser          string
	Queue            *int
	Status           *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Session          model.Sessions
	HealthFacilities HealthFacilitiesCustomUser
}

type HealthFacilitiesCustomUser struct {
	ID        string
	Email     string
	PhoneNum  string
	Name      string
	Image     *string
	CreatedAt time.Time
	UpdatedAt time.Time
	Address   *model.Addresses
}

type AgeUser struct {
	BirthDate time.Time
	Age       int
}

type UserNearbyHealth struct {
	User             UserProfile
	HealthFacilities []HealthResponse
}

type UpdateUser struct {
	Fullname  string
	NikUser   string
	Email     string
	Gender    string
	PhoneNum  string
	BirthDate time.Time
}
