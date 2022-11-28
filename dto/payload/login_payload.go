package payload

type Login struct {
	Email    string `json:"email" gorm:"not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}