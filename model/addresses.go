package model

import "gorm.io/gorm"

type Addresses struct {
	ID             string
	CurrentAddress string
	District       string
	City           string
	Province       string
	Longitude      float64
	Latitude       float64
	CreatedAt      gorm.CreatedAt
	UpdatedAt      gorm.UpdatedAt
	DeletedAt      gorm.DeletedAt
}
