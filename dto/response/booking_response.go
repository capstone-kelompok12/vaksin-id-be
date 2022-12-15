package response

import (
	"time"
	"vaksin-id-be/model"
)

type BookingResponse struct {
	ID        string
	IdSession string
	Queue     *int
	Status    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	Session   model.Sessions
	History   []model.VaccineHistories
}
type BookingResponseCustom struct {
	ID        string
	IdSession string
	Queue     *int
	Status    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	Session   BookingSessionCustom
	History   []BookingHistoryCustom
}

type BookingFindQueue struct {
	ID        string
	IdSession string
	Queue     int
	Status    string
}

type BookingSessionCustom struct {
	ID string
	// IdHealthFacilities string
	IdVaccine    string
	SessionName  string
	Capacity     int
	Dose         int
	Date         time.Time
	IsClose      bool
	StartSession string
	EndSession   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Vaccine      model.Vaccines
}

type BookingHistoryCustom struct {
	ID         string
	IdBooking  string
	NikUser    string
	IdSameBook string
	Status     *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       *model.Users
}
