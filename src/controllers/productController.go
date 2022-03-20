package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sort"
	"strconv"
	"strings"
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

	go database.ClearCache("productsFrontend", "productsBackend")

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
	database.AqlNoReturn(dbQuery)

	//database.Cache.Del(context.Background(), "productsFrontend", "productsBackend")
	//go clearCache("productsFrontend", "productsBackend")
	go database.ClearCache("productsFrontend", "productsBackend")

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")

	dbQuery := fmt.Sprintf("REMOVE \"%s\" IN Products", id)
	println(dbQuery)
	database.AqlNoReturn(dbQuery)
	go database.ClearCache("productsFrontend", "productsBackend")

	return c.JSON(fiber.Map{
		"message": "Product deleted.",
	})
}

func ProductsFrontend(c *fiber.Ctx) error {

	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "productsFrontend").Result()

	if result == "" || err != nil {
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

func ProductsBackend(c *fiber.Ctx) error {

	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "productsBackend").Result()

	if result == "" || err != nil {

		products = database.ReturnArrayOfObject("FOR r in Products RETURN r", products)

		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}

		err2 := database.Cache.Set(ctx, "productsBackend", bytes, time.Minute*30).Err()
		if err2 != nil {
			panic(err2)
		}

	} else {
		err := json.Unmarshal([]byte(result), &products)
		if err != nil {
			return err
		}
	}

	var filteredProducts []models.Product

	if s := c.Query("s"); s != "" {

		for _, product := range products {
			if strings.Contains("Products/"+product.Id, s) || strings.Contains(product.Title, strings.Title(s)) {
				filteredProducts = append(filteredProducts, product)
			}

		}

	} else {
		filteredProducts = products
	}

	if sortQuery := c.Query("sort"); sortQuery != "" {
		sortLower := strings.ToLower(sortQuery)
		if sortLower == "asc" {
			sort.Slice(filteredProducts, func(i, j int) bool {
				return filteredProducts[i].Price < filteredProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(filteredProducts, func(i, j int) bool {
				return filteredProducts[i].Price > filteredProducts[j].Price
			})
		}
	}

	var total = len(filteredProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9

	var data []models.Product

	if page < 1 {
		page = 1
	}

	if total <= page*perPage && total >= (page-1)*perPage {
		data = filteredProducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = filteredProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}

	return c.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})

}
