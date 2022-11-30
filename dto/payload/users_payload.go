package payload

type RegisterUser struct {
	Fullname  string `json:"fullname" gorm:"size:255;not null" validate:"required" example:"test"`
	NikUser   string `json:"nik" gorm:"varchar:16;not null" validate:"required,min=16,max=16" example:"test"`
	Email     string `json:"email" gorm:"size:100;not null" validate:"required,email"`
	Gender    string `json:"gender" gorm:"size:1;not null" validate:"required"`
	Password  string `json:"password" gorm:"size:100;not null" validate:"required,min=6"`
	PhoneNum  string `json:"phonenum" gorm:"size:15;not null" validate:"required,min=10,max=15"`
	BirthDate string `json:"birthdate" gorm:"not null" validate:"required"`
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
