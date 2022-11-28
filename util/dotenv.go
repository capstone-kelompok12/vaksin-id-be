package util

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ProcessEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Error loading env file!")
		panic(err)
	}
}
