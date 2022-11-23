package main

import (
	"os"
	"vaksin-id-be/routes"
)

func main() {

	route := routes.Init()

	port := ":" + os.Getenv("PORT")

	route.Logger.Fatal(route.Start(port))

}
