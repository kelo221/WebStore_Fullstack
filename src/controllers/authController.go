package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
	"webstore/src/database"
	"webstore/src/middlewares"
	"webstore/src/models"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["email"] == "" || data["password"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Missing email or password!",
		})
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Avatar:    "/img/users/noImage.jpg",
		IsAdmin:   !strings.Contains(c.Path(), "api/user"),
	}

	if user.Email == "user" {
		user.Avatar = "/img/users/user.png"
	}

	if user.Email == "admin" {
		user.Avatar = "/img/users/admin.png"
	}

	user.SetPassword(data["password"])

	database.PushUser(&user)

	return c.JSON(fiber.Map{"message": "Registered!"})
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

	var privilege models.UserType

	if strings.Contains(c.Path(), "api/user") {
		privilege = models.Regular
	} else {
		privilege = models.Admin
	}

	token, err := middlewares.GenerateJWT(user.Id, privilege)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "None",
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
	user.Password = nil

	return c.JSON(user)

}

func LogOut(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "None",
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged out.",
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
