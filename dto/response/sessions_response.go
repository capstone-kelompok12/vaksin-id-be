package response

import (
	"time"
	"vaksin-id-be/model"
)

type SessionsResponse struct {
	ID           string
	SessionName  string
	Capacity     int
	IsClose      bool
	Dose         int
	StartSession string
	EndSession   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Booking      []model.BookingSessions
}
type SessionsUpdate struct {
	ID           string
	SessionName  string
	Capacity     int
	IsClose      bool
	StartSession string
	EndSession   string
}
