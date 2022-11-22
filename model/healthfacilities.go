package model

import (
	"time"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	ID        string
	IDAddress string
	Email     string
	Password  string
	PhoneNum  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
