package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
	"webstore/src/database"
	"webstore/src/models"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product
	//Added generics, keeping this incase something breaks
	//products = database.AqlReturnProducts("FOR r in Products RETURN r")
	products = database.ReturnArrayOfObject("FOR r in Products RETURN r", products)

	return c.JSON(products)
}

func Product(c *fiber.Ctx) error {

	var product models.Product
	id := c.Params("id")

	dbQuery := fmt.Sprintf("FOR r IN Products FILTER r._key == \"%s\" RETURN r", id)
	println(dbQuery)
	//product = database.AqlReturnProduct(dbQuery)
	product = database.ReturnObject(dbQuery, product)

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

func ProductsFrontend(c *fiber.Ctx) error {

	var products []models.Product
	var ctx = context.Background()

	result, _ := database.Cache.Get(ctx, "productsFrontend").Result()

	if result == "" {
		fmt.Println("cache not found")

		products = database.ReturnArrayOfObject("FOR r in Products RETURN r", products)

		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}

		err2 := database.Cache.Set(ctx, "productsFrontend", bytes, time.Minute*30).Err()
		if err2 != nil {
			panic(err2)
		}
	} else {
		fmt.Println("found cache")
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)

}
