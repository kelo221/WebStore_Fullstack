package routes

import (
	"github.com/gofiber/fiber/v2"
	"webstore/src/controllers"
	"webstore/src/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("api")

	admin := api.Group("admin")

	//admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuth := admin.Use(middlewares.IsAuthAdmin)

	adminAuth.Get("", controllers.User)
	adminAuth.Post("logout", controllers.LogOut)
	adminAuth.Put("info", controllers.UpdateInfo)
	adminAuth.Put("password", controllers.UpdatePassword)
	adminAuth.Get("admins", controllers.Admins)

	adminAuth.Get("products", controllers.Products)
	adminAuth.Get("products/:id", controllers.Product)

	adminAuth.Post("products", controllers.CreateProducts)
	adminAuth.Put("products/:id", controllers.UpdateProduct)
	adminAuth.Delete("products/:id", controllers.DeleteProduct)

	adminAuth.Get("orders", controllers.Orders)

	userRoute := api.Group("user")
	userRoute.Post("register", controllers.Register)
	userRoute.Post("login", controllers.Login)

	userRouteAuth := userRoute.Use(middlewares.IsAuth)
	userRouteAuth.Get("info", controllers.User)
	userRouteAuth.Post("logout", controllers.LogOut)
	userRouteAuth.Put("info", controllers.UpdateInfo)
	userRouteAuth.Put("password", controllers.UpdatePassword)

	userRouteAuth.Get("orders", controllers.OrdersLimited)

	api.Get("products/frontend", controllers.ProductsFrontend)
	api.Get("products/backend", controllers.ProductsBackend)

	checkOut := api.Group("checkout")
	checkOut.Put("orders", controllers.CreateOrder)

}
