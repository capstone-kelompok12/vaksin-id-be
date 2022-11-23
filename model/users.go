package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	NIK          int       `gorm:"type:bigint(16);primary_key"`
	Username     string    `gorm:"type:varchar(255);unique"`
	Password     string    `gorm:"type:varchar(255)"`
	Email        string    `gorm:"type:varchar(255);unique"`
	Fullname     string    `gorm:"type:varchar(255)"`
	PhoneNum     string    `gorm:"type:varchar(255)"`
	VaccineCount int       `gorm:"type:int(11)"`
	BirthDate    time.Time `gorm:"type:date"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Address      *Addresses     `gorm:"foreignKey:NikUser"` // belong to relationship
}
