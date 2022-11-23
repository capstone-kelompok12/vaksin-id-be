package payload

import "time"

type Register struct {
	NikUser   string `json:nik`
	Username  string `json:"username" gorm:"size:255;not null" validate:"required,min=4"`
	Password  string `json:"password" gorm:"size:100;not null" validate:"required,min=6"`
	Email     string `json:"email" gorm:"size:100;not null" validate:"required,email"`
	Fullname  string `json:"fullname" gorm:"size:255;not null" validate:"required"`
	PhoneNum  string
	BirthDate time.Time
}
