package model

import (
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	ID                 string
	IDHealthFacilities string
	Email              string
	Password           string
	PhoneNum           string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
