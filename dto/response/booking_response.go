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

// type BookingResponse struct {
// 	ID        string
// 	NikUser   string
// 	IdSession string
// }
