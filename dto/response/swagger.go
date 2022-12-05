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

type HealthFacilitiesSwagger struct {
	ID        string  `gorm:"type:varchar(255);primary_key"`
	Email     string  `gorm:"type:varchar(255)"`
	PhoneNum  string  `gorm:"type:varchar(255)"`
	Name      string  `gorm:"type:varchar(255)"`
	Image     *string `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time         `gorm:"index"`
	Address   *AddressesSwagger `gorm:"foreignKey:IdHealthFacilities"`
}

type AddressesSwagger struct {
	ID                 string  `gorm:"type:varchar(255);primary_key"`
	IdHealthFacilities *string `gorm:"type:varchar(255)"`
	NikUser            *string `gorm:"type:varchar(16)"`
	CurrentAddress     string  `gorm:"type:longtext"`
	District           string  `gorm:"type:varchar(255)"`
	City               string  `gorm:"type:varchar(255)"`
	Province           string  `gorm:"type:varchar(255)"`
	Latitude           float64 `gorm:"type:numeric(11,7)"`
	Longitude          float64 `gorm:"type:numeric(11,7)"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time `gorm:"index"`
}

type Response[T any] struct {
	Message string `json:"message" example:"success"`
	Data    any    `json:"data"`
}
type ResponseError[T any] struct {
	Message string `json:"message" example:"error"`
	Data    string `json:"data" example:""`
	Error   bool   `json:"error"`
}
type ResponseDelete[T any] struct {
	Message string `json:"message" example:"success deleted"`
	Data    string `json:"data" example:""`
}
