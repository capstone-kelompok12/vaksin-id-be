package main

import (
	"os"
	"vaksin-id-be/routes"
)

func main() {

	route := routes.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routePort := ":" + port

	route.Logger.Fatal(route.Start(routePort))

}
