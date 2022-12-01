package response

import "time"

type AdminResponse struct {
	ID                 string
	IdHealthFacilities string
	Email              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
