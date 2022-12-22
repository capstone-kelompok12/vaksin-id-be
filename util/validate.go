package util

import (
	"github.com/go-playground/validator/v10"
)

func Validate(payloads interface{}) error {
	validate := validator.New()
	return validate.Struct(payloads)
}
