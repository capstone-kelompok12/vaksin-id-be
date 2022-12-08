package model

import (
	"time"
)

type Sessions struct {
	ID                 string `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string `gorm:"type:varchar(255)"`
	SessionName        string `gorm:"type:varchar(255)"`
	Capacity           int    `gorm:"type:int(11)"`
	Dose               string `gorm:"type:varchar(2)"`
	IsClose            bool   `gorm:"type:boolean"`
	StartSession       string `gorm:"type:varchar(5)"`
	EndSession         string `gorm:"type:varchar(5)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
