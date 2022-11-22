package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	NIK          int
	IDAddress    string
	Username     string
	Password     string
	Email        string
	Fullname     string
	PhoneNum     string
	VaccineCount int
	BirthDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
