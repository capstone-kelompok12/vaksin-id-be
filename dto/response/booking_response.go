package response

import "time"

type BookingResponse struct {
	ID        string
	NikUser   string
	IdSession string
	Queue     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
