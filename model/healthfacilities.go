package model

import (
	"time"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	ID        string  `gorm:"type:varchar(255);primary_key"`
	Email     string  `gorm:"type:varchar(255)"`
	PhoneNum  string  `gorm:"type:varchar(255)"`
	Name      string  `gorm:"type:varchar(255)"`
	Image     *string `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Address   *Addresses     `gorm:"foreignKey:IdHealthFacilities"` // has one to relationship
	Vaccine   []Vaccines     `gorm:"foreignKey:IdHealthFacilities"` // has many relationship
}
