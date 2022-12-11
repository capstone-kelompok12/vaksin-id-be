package model

import (
	"time"
)

type Sessions struct {
	ID                 string `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string `gorm:"type:varchar(255)"`
	SessionName        string `gorm:"type:varchar(255)"`
	Capacity           int    `gorm:"type:int(11)"`
	Dose               int    `gorm:"type:int(1)"`
	IsClose            bool   `gorm:"type:boolean"`
	StartSession       string `gorm:"type:varchar(5)"`
	EndSession         string `gorm:"type:varchar(5)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Booking            []BookingSessions `gorm:"foreignKey:IdSession"`
}
