package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"webstore/src/commands"
	"webstore/src/database"
	"webstore/src/routes"
)

///		TODO
/// 	1.	Check if email already exists
///		2.	Connect user orders

func main() {

	database.ConnectDB()
	database.SetupRedis()
	database.SetupCacheChannel()

	app := fiber.New()

	commands.Populate()

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://127.0.0.1:8000, http://127.0.0.1:8000,  http://localhost:8000,  http://localhost:3000,",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	//go database.ClearCache("productsFrontend", "productsBackend")

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
