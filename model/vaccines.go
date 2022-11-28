package model

import "time"

type Vaccines struct {
	ID                 string `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string `gorm:"type:varchar(255)"`
	Name               string `gorm:"type:varchar(255)"`
	Stock              int    `gorm:"type:int(11)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
