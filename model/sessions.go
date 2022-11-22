package model

import (
	"time"
)

type Sessions struct {
	ID                 string
	IDHealthFacilities string
	SessionName        string
	Capacity           int
	SessionStatus      bool
	StartSession       time.Time
	EndSession         time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
