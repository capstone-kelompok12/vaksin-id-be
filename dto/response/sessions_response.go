package response

import "time"

type SessionsResponse struct {
	ID            string
	SessionName   string
	Capacity      int
	SessionStatus bool
	StartSession  time.Time
	EndSession    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
