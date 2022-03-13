package controllers

import (
	"ambassor/src/database"
	"ambassor/src/models"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product
	products = database.AqlReturnProducts("FOR r in Products RETURN r")

	return c.JSON(products)
}

func Product(c *fiber.Ctx) error {

	var product models.Product
	id := c.Params("id")

	dbQuery := fmt.Sprintf("FOR r IN Products FILTER r._key == \"%s\" RETURN r", id)
	println(dbQuery)
	product = database.AqlReturnProduct(dbQuery)

	return c.JSON(product)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	database.PushProduct(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	if err := c.BodyParser(&product); err != nil {
		println("parsing error")
		return err
	}

	newProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	dbQuery := fmt.Sprintf("UPDATE \"%s\" WITH %s IN Products", id, newProduct)
	println(dbQuery)
	database.AqlNoReturn(dbQuery)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	dbQuery := fmt.Sprintf("REMOVE \"%s\" IN Products", id)
	println(dbQuery)
	database.AqlNoReturn(dbQuery)

	return c.JSON(fiber.Map{
		"message": "Product deleted.",
	})
}
