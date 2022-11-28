package util

import (
	"vaksin-id-be/dto/payload"

	"github.com/go-playground/validator/v10"
)

func ValidateRegister(payloads payload.RegisterUser) error {
	validate := validator.New()
	return validate.Struct(payloads)
}
