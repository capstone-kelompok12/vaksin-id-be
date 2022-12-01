package response

import "time"

type AdminsResponse struct {
	ID                 string
	IdHealthFacilities string
	Email              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
