package model

import (
	"time"

	"gorm.io/gorm"
)

type BookingSessions struct {
	ID        string `gorm:"type:varchar(255);primary_key"`
	NikUser   string `gorm:"type:varchar(16)"`
	IdSession string `gorm:"type:varchar(255)"`
	Queue     int    `gorm:"type:int(11)"`
	Status    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      *Users         `gorm:"foreignKey:NikUser"` // belong to relationship
	// Session   *Sessions          `gorm:"foreignKey:IdSession"` // belong to relationship
	Histroy []VaccineHistories `gorm:"foreignKey:IdBooking"` // has many relationship
}
