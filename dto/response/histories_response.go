package response

import (
	"time"
	"vaksin-id-be/model"
)

type HistoryResponse struct {
	ID         string
	IdBooking  string
	NikUser    string
	IdSameBook string
	Status     *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       model.Users
}
