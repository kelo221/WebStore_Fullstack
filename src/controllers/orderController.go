package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"webstore/src/database"
	"webstore/src/middlewares"
	"webstore/src/models"
)

func Orders(c *fiber.Ctx) error {

	var orders []models.ShoppingCart
	orders = database.ReturnArrayOfObject("FOR r in Orders RETURN r", orders)

	for i, orderObject := range orders {
		orders[i].Name = orderObject.FullName()
		orders[i].Total = orderObject.CalculateTotal()
	}

	return c.JSON(orders)
}

func OrdersLimited(c *fiber.Ctx) error {

	id, _ := middlewares.GetUserID(c)

	var result models.ShoppingCart

	dbQuery := fmt.Sprintf("FOR r in Orders FILTER r.UserId == \"%s\" RETURN r", id)
	fmt.Printf("%s \n", dbQuery)
	user := database.ReturnObject(dbQuery, result)

	return c.JSON(user)
}

/*
type CreateOrderRequest struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Name      string `json:"Name,omitempty"`

	Address string `json:"Address"`
	City    string `json:"City"`
	Country string `json:"Country"`
	Zip     string `json:"Zip"`

	Total    float64          `json:"Total,omitempty"`
	Products []map[string]int `json:"Products"`
}*/

func CreateOrder(c *fiber.Ctx) error {

	var request models.ShoppingCart

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	/*order := models.ShoppingCart{
		Id:            "",
		TransactionId: "",
		UserId:        request.FirstName,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Email:         request.Email,
		Name:          request.Name,
		Address:       request.Address,
		City:          request.City,
		Country:       request.Country,
		Zip:           request.Zip,
		Complete:      false,
		Total:         0,
		OrderItems:    nil,
	}*/

	database.PushShoppingList(&request)
	return c.JSON(request)
}
