package controllers

import (
	"ambassor/src/database"
	"ambassor/src/middlewares"
	"ambassor/src/models"
	"encoding/json"
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

	//fmt.Printf("%+v\n", user)

	return c.JSON(fiber.Map{"message": "hello"})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	dbQuery := fmt.Sprintf("FOR r IN Users FILTER r.Email == \"%s\" RETURN r", data["email"])
	user := database.AqlReturnUser(dbQuery)

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

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	cookie := fiber.Cookie{
		Name:        "jwt",
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

func User(c *fiber.Ctx) error {

	id, _ := middlewares.GetUserID(c)

	dbQuery := fmt.Sprintf("FOR r IN Users FILTER r._id == \"%s\" RETURN r", id)
	user := database.AqlReturnUser(dbQuery)

	return c.JSON(user)

}

func LogOut(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:        "jwt",
		Value:       "",
		Path:        "",
		Domain:      "",
		MaxAge:      0,
		Expires:     time.Now().Add(-time.Hour),
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

func UpdateInfo(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		println("parsing error")
		return err
	}

	id, _ := middlewares.GetUserID(c)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	newUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	println(newUser)

	dbQuery := fmt.Sprintf("UPDATE DOCUMENT(\"%s\") WITH %s IN Users", id, newUser)
	println(dbQuery)
	database.AqlNoReturn(dbQuery)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {

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

	id, _ := middlewares.GetUserID(c)

	var user models.User

	user.SetPassword(data["password"])

	newUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	dbQuery := fmt.Sprintf("UPDATE DOCUMENT(\"%s\") WITH %s IN Users", id, newUser)
	database.AqlNoReturn(dbQuery)

	return c.JSON(fiber.Map{
		"message": "Success.",
	})
}
