package routes

import (
	"ambassor/src/controllers"
	"ambassor/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuth := admin.Use(middlewares.IsAuth)
	adminAuth.Get("user", controllers.User)
	adminAuth.Post("logout", controllers.LogOut)
	adminAuth.Put("user/info", controllers.UpdateInfo)
	adminAuth.Put("user/password", controllers.UpdatePassword)

}
