package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.ShoppingCart

	orders = database.AqlReturnOrders("FOR r in Orders RETURN r")

	return c.JSON(orders)
}
