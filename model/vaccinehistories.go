package model

import "time"

type VaccineHistories struct {
	ID        string `gorm:"type:varchar(255);primary_key"`
	IdBooking string `gorm:"type:varchar(255)"`
	Status    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
}
