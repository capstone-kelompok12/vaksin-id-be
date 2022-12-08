package response

import "time"

type SessionsResponse struct {
	ID           string
	SessionName  string
	Capacity     int
	IsClose      bool
	StartSession string
	EndSession   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type SessionsUpdate struct {
	ID           string
	SessionName  string
	Capacity     int
	IsClose      bool
	StartSession string
	EndSession   string
}
