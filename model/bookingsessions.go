package model

import (
	"time"

	"gorm.io/gorm"
)

type BookingSessions struct {
	ID        string
	NIKUser   int
	IDSession string
	Queue     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
