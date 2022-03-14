package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"github.com/gofiber/fiber/v2"
)

func Admins(c *fiber.Ctx) error {
	var users []models.User
	users = database.AqlReturnUsers("FOR r in Users FILTER r.IsAdmin == true RETURN r")

	return c.JSON(users)
}