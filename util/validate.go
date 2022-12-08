package util

import (
	"vaksin-id-be/dto/payload"

	"github.com/go-playground/validator/v10"
)

func ValidateRegister(payloads payload.RegisterUser) error {
	validate := validator.New()
	return validate.Struct(payloads)
}

func ValidateHealthFacilities(payloads payload.HealthFacilities) error {
	validate := validator.New()
	return validate.Struct(payloads)
}

func ValidateVaccine(payloads payload.VaccinesPayload) error {
	validate := validator.New()
	return validate.Struct(payloads)
}

func ValidateSession(payloads payload.SessionsPayload) error {
	validate := validator.New()
	return validate.Struct(payloads)
}
