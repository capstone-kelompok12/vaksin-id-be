package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	NIK          string    `gorm:"type:varchar(16);primary_key"`
	Email        string    `gorm:"type:varchar(255);unique"`
	Password     string    `gorm:"type:varchar(255)"`
	Fullname     string    `gorm:"type:varchar(255)"`
	PhoneNum     string    `gorm:"type:varchar(255)"`
	ProfileImage *string   `gorm:"type:longtext"`
	Gender       string    `gorm:"type:enum('P', 'L')"`
	VaccineCount int       `gorm:"type:int(11)"`
	BirthDate    time.Time `gorm:"type:date"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *gorm.DeletedAt `gorm:"index"`
	Address      *Addresses      `gorm:"foreignKey:NikUser"` // has one to relationship
}
