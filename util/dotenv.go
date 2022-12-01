package util

import (
	"github.com/joho/godotenv"
)

func ProcessEnv() {
	godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Error("Error loading env file!")
	// 	panic(err)
	// }
}
