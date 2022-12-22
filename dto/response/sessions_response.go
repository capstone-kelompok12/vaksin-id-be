package response

import (
	"time"
	"vaksin-id-be/model"
)

type SessionsResponse struct {
	ID           string
	IdVaccine    string
	SessionName  string
	Capacity     int
	CapacityLeft int
	IsClose      bool
	Dose         int
	Date         time.Time
	StartSession string
	EndSession   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Vaccine      model.Vaccines
	Booking      []BookingInSession
}

type BookingInSession struct {
	ID        string
	IdSession string
	NikUser   string
	Queue     int
	Status    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *model.Users
	History   SessionCustomHistory
}

type SessionCustomHistory struct {
	ID         string
	IdBooking  string
	NikUser    string
	IdSameBook string
	Status     string
}

type SessionUserCustom struct {
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

type SessionsUpdate struct {
	ID           string
	IdVaccine    string
	SessionName  string
	Capacity     int
	CapacityLeft int
	Dose         int
	Date         time.Time
	IsClose      bool
	StartSession string
	EndSession   string
}

type SessionSumCap struct {
	ID            string
	TotalCapacity int
}
