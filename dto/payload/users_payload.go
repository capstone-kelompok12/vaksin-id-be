package payload

import "time"

type RegisterUser struct {
	Fullname  string    `json:"fullname" gorm:"size:255;not null" validate:"required"`
	NikUser   string    `json:"nik" gorm:"varchar:16;not null" validate:"required,min=16,max=16"`
	Email     string    `json:"email" gorm:"size:100;not null" validate:"required,email"`
	Gender    string    `json:"username" gorm:"size:1;not null" validate:"required"`
	Password  string    `json:"password" gorm:"size:100;not null" validate:"required,min=6"`
	PhoneNum  string    `json:"phonenum" gorm:"size:15;not null" validate:"required,min=10,max=15"`
	BirthDate time.Time `json:"birthdate" gorm:"not null" validate:"required"`
}

type UpdateUser struct {
	Fullname  string    `json:"fullname" gorm:"size:255"`
	Email     string    `json:"email" gorm:"size:100" validate:"email"`
	Gender    string    `json:"username" gorm:"size:255" validate:"min=4"`
	PhoneNum  string    `json:"phonenum" gorm:"size:15" validate:"min=10,max=15"`
	BirthDate time.Time `json:"birthdate"`
}
