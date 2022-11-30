package payload

type HealthFacilities struct {
	Email          string  `json:"email" validate:"required,email"`
	PhoneNum       string  `json:"phonenum"`
	Name           string  `json:"name"`
	Image          *string `json:"image"`
	CurrentAddress string  `json:"current_address"`
	District       string  `json:"district"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Latitude       float64 `json:"latitude" gorm:"type:numeric(11,7)"`
	Longitude      float64 `json:"longitude" gorm:"type:numeric(11,7)"`
}
