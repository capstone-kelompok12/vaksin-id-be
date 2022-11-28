package model

import (
	"time"
)

type Sessions struct {
	ID                 string `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string `gorm:"type:varchar(255)"`
	SessionName        string `gorm:"type:varchar(255)"`
	Capacity           int    `gorm:"type:int(11)"`
	SessionStatus      bool   `gorm:"type:boolean"`
	StartSession       time.Time
	EndSession         time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
