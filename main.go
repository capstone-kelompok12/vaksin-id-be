package main

import (
	"os"
	"vaksin-id-be/routes"
)

// @title VAKSIN-ID API
// @version 1.0
// @description This is a Booking Vaccine API for manage Booking
// @description Capstone Project Kelompok 12
// @BasePath  /api/v1
// @schemes http https
func main() {

	route := routes.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routePort := ":" + port

	route.Logger.Fatal(route.Start(routePort))
}
