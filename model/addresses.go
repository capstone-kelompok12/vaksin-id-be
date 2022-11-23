package model

import (
	"time"

	"gorm.io/gorm"
)

type Addresses struct {
	ID                 string  `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string  `gorm:"type:varchar(255)"`
	NikUser            int     `gorm:"type:bigint(16)"`
	CurrentAddress     string  `gorm:"type:longtext"`
	District           string  `gorm:"type:varchar(255)"`
	City               string  `gorm:"type:varchar(255)"`
	Province           string  `gorm:"type:varchar(255)"`
	Longitude          float64 `gorm:"type:numeric(11,7)"`
	Latitude           float64 `gorm:"type:numeric(11,7)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
