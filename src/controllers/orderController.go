package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {

	var orders []models.ShoppingCart
	//orders = database.AqlReturnOrders("FOR r in Orders RETURN r")
	orders = database.ReturnArrayOfObject("FOR r in Orders RETURN r", orders)

	for i, orderObject := range orders {
		orders[i].Name = orderObject.FullName()
		orders[i].Total = orderObject.CalculateTotal()
	}

	return c.JSON(orders)
}
