package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"webstore/src/database"
	"webstore/src/routes"
)

func main() {

	database.ConnectDB()
	database.SetupRedis()
	database.SetupCacheChannel()

	app := fiber.New()

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://127.0.0.1:8000,  http://localhost:8000,  http://localhost:3000,",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
