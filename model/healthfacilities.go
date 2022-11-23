package model

import (
	"time"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	ID        string `gorm:"type:varchar(255);primary_key"`
	Email     string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(255)"`
	PhoneNum  string `gorm:"type:varchar(255)"`
	Name      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Address   *Addresses     `gorm:"foreignKey:IdHealthFacilities"` // belong to relationship
	Session   []Sessions     `gorm:"foreignKey:IdHealthFacilities"` // has many relationship
	Vaccine   []Vaccines     `gorm:"foreignKey:IdHealthFacilities"` // has many relationship
}
