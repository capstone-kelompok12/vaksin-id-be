package util

import (
	"vaksin-id-be/dto/payload"

	"github.com/go-playground/validator/v10"
)

func ValidateRegister(payload payload.Register) error {
	validate := validator.New()
	return validate.Struct(payload)
}
