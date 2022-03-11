package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
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

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		IsAdmin:   false,
	}

	user.SetPassword(data["password"])

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

	if user.Email == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})

	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	payload := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Id:        "",
		IssuedAt:  0,
		Issuer:    "",
		NotBefore: 0,
		Subject:   user.Id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("test"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	cookie := fiber.Cookie{
		Name:        "JWT",
		Value:       token,
		Path:        "",
		Domain:      "",
		MaxAge:      0,
		Expires:     time.Now().Add(time.Hour * 24),
		Secure:      false,
		HTTPOnly:    true,
		SameSite:    "",
		SessionOnly: false,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success.",
	})
}
