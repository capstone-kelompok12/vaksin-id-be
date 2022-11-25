package util

import (
	"github.com/go-playground/validator/v10"
)

func Validate(payload any) error {
	validate := validator.New()
	return validate.Struct(payload)
}
