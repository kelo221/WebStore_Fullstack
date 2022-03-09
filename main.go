package main

import (
	"ambassor/src/database"
	"ambassor/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
