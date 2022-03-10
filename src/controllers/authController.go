package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
		IsAdmin:   false,
	}

	database.PushUser(&user)

	return c.JSON(fiber.Map{"message": "hello"})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	dbQuery := fmt.Sprintf("FOR r IN Users FILTER r.Email == \"%s\" RETURN r", data["email"])
	println("query: ", dbQuery)
	user := database.AqlJSON(dbQuery)

	println("returned from db: ", user.Email)

	if user.Email == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})

	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte((data["password"]))); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	return c.JSON(user)
}
