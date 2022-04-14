package main

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
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

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://127.0.0.1:8000,  https://localhost:8000,  https://localhost:3000, https://127.0.0.1:8000,  https://localhost:8000,  http://localhost:3001,",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(c.Protocol()) // => https
	})

	routes.Setup(app)

	// Create tls certificate
	cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Create custom listener
	ln, err := tls.Listen("tcp", "127.0.0.1:8000", config)
	if err != nil {
		panic(err)
	}

	// Start server with https/ssl enabled on http://localhost:443
	log.Fatal(app.Listener(ln))
}
