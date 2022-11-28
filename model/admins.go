package model

import (
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	ID                 string `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities string `gorm:"type:varchar(255)"`
	Email              string `gorm:"type:varchar(255)"`
	Password           string `gorm:"type:varchar(255)"`
	PhoneNum           string `gorm:"type:varchar(255)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt    `gorm:"index"`
	HealthFacility     *HealthFacilities `gorm:"foreignKey:IdHealthFacilities"` // belong to relationship
}
