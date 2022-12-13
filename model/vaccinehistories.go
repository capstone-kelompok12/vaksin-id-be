package model

import "time"

type VaccineHistories struct {
	ID         string  `gorm:"type:varchar(255);primary_key"`
	IdBooking  string  `gorm:"type:varchar(255)"`
	NikUser    string  `gorm:"type:varchar(16)"`
	IdSameBook string  `gorm:"type:varchar(255)"`
	Status     *string `gorm:"type:varchar(255)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       *Users           `gorm:"foreignKey:NikUser"`   // belong to relationship
	Booking    *BookingSessions `gorm:"foreignKey:IdBooking"` // belong to relationship
}
