package routes

import (
	"github.com/gofiber/fiber/v2"
	"webstore/src/controllers"
	"webstore/src/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuth := admin.Use(middlewares.IsAuth)

	adminAuth.Get("", controllers.User)
	adminAuth.Post("logout", controllers.LogOut)
	adminAuth.Put("user/info", controllers.UpdateInfo)
	adminAuth.Put("user/password", controllers.UpdatePassword)
	adminAuth.Get("admins", controllers.Admins)

	adminAuth.Get("products", controllers.Products)
	adminAuth.Get("products/:id", controllers.Product)
	adminAuth.Post("products", controllers.CreateProducts)
	adminAuth.Post("products/:id", controllers.UpdateProduct)
	adminAuth.Delete("products/:id", controllers.DeleteProduct)
	adminAuth.Get("orders", controllers.Orders)

	userRoute := api.Group("user")
	userRoute.Post("register", controllers.Register)
	userRoute.Post("login", controllers.Login)

	userRouteAuth := userRoute.Use(middlewares.IsAuth)
	userRouteAuth.Get("", controllers.User)
	userRouteAuth.Post("logout", controllers.LogOut)
	userRouteAuth.Put("user/info", controllers.UpdateInfo)
	userRouteAuth.Put("user/password", controllers.UpdatePassword)

	userRouteAuth.Get("products/frontend", controllers.ProductsFrontend)

	//TODO personal orders tied to user id, maybe

}
