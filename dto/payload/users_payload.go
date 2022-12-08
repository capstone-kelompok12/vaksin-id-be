package payload

type RegisterUser struct {
	Fullname  string `json:"fullname" gorm:"size:255;not null" validate:"required" example:"user"`
	NikUser   string `json:"nik" gorm:"varchar:16;not null" validate:"required,min=16,max=16" example:"1234567898765432"`
	Email     string `json:"email" gorm:"size:100;not null" validate:"required,email" example:"user@gmail.com"`
	Gender    string `json:"gender" gorm:"size:1;not null" validate:"required" example:"L"`
	Password  string `json:"password" gorm:"size:100;not null" validate:"required,min=6" example:"user123"`
	PhoneNum  string `json:"phonenum" gorm:"size:15;not null" validate:"required,min=10,max=15" example:"081234567890"`
	BirthDate string `json:"birthdate" gorm:"not null" validate:"required" example:"2001-05-25"`
}

type UpdateUser struct {
	Fullname     string `json:"fullname" gorm:"size:255"`
	Email        string `json:"email" gorm:"size:100"`
	Gender       string `json:"username" gorm:"size:1"`
	PhoneNum     string `json:"phonenum" gorm:"size:15"`
	ProfileImage string `json:"profileimage"`
	BirthDate    string `json:"birthdate"`
}

type UpdateAddress struct {
	CurrentAddress string  `json:"current_address"`
	District       string  `json:"district"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Latitude       float64 `json:"latitude" gorm:"type:numeric(11,7)"`
	Longitude      float64 `json:"longitude" gorm:"type:numeric(11,7)"`
}

type NearbyHealth struct {
	Latitude  float64 `json:"latitude" gorm:"type:numeric(11,7)"`
	Longitude float64 `json:"longitude" gorm:"type:numeric(11,7)"`
}
