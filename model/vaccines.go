package model

import "time"

type Vaccines struct {
	ID                 string
	IDHealthFacilities string
	Name               string
	Stock              int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
