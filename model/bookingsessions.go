package model

import (
	"time"

	"gorm.io/gorm"
)

type BookingSessions struct {
	ID        string  `gorm:"type:varchar(255);primary_key"`
	IdSession string  `gorm:"type:varchar(255)"`
	Queue     *int    `gorm:"type:int(11)"`
	Status    *string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt     `gorm:"index"`
	Session   *Sessions          `gorm:"foreignKey:IdSession"` // belong to relationship
	History   []VaccineHistories `gorm:"foreignKey:IdBooking"` // has many relationship
}
