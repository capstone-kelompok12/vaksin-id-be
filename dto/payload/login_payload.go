package payload

type Login struct {
	Email    string `json:"email" gorm:"not null" validate:"required,email" example:"user@gmail.com"`
	Password string `json:"password" gorm:"not null" validate:"required" example:"user123"`
}
