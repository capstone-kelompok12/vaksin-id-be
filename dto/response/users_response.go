package response

import (
	"time"
)

type UserProfile struct {
	NIK          string
	Email        string
	Fullname     string
	PhoneNum     string
	Gender       string
	VaccineCount int
	Age          int
}

type AgeUser struct {
	BirthDate time.Time
	Age       int
}
