package response

import "time"

type UserAddresses struct {
	ID                 string  `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities *string `gorm:"type:varchar(255)"`
	NikUser            *string `gorm:"type:varchar(16)"`
	CurrentAddress     string  `gorm:"type:longtext"`
	District           string  `gorm:"type:varchar(255)"`
	City               string  `gorm:"type:varchar(255)"`
	Province           string  `gorm:"type:varchar(255)"`
	Longitude          float64 `gorm:"type:numeric(11,7)"`
	Latitude           float64 `gorm:"type:numeric(11,7)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time `gorm:"index"`
}

type Response[T any] struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
