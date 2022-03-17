package controllers

import (
	"github.com/gofiber/fiber/v2"
	"webstore/src/database"
	"webstore/src/models"
)

func Admins(c *fiber.Ctx) error {
	var users []models.User
	//users = database.AqlReturnUsers("FOR r in Users FILTER r.IsAdmin == true RETURN r")
	users = database.ReturnArrayOfObject("FOR r in Users FILTER r.IsAdmin == true RETURN r", users)

	return c.JSON(users)
}
