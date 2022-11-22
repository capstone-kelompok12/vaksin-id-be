package model

import "time"

type VaccineHistories struct {
	ID        string
	IDBooking string
	Status    string
	CreatedAt time.Time
}
