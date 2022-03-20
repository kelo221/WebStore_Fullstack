package controllers

import (
	"github.com/gofiber/fiber/v2"
	"webstore/src/database"
	"webstore/src/models"
)

func Orders(c *fiber.Ctx) error {

	var orders []models.ShoppingCart
	//orders = database.AqlReturnOrders("FOR r in Orders RETURN r")
	orders = database.ReturnArrayOfObject("FOR r in Orders RETURN r", orders)

	for i, orderObject := range orders {
		orders[i].Name = orderObject.FullName()
		orders[i].Total = orderObject.CalculateTotal()
	}

	return c.JSON(orders)
}

type CreateOrderRequest struct {
	Code string `json:"Code"`

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
}

func CreateOrder(c *fiber.Ctx) error {

	/*	var request = CreateOrderRequest

		err := c.BodyParser(&request)
		if err != nil {
			return err
		}

		order := models.ShoppingCart{
			Id:            "",
			TransactionId: "",
			UserId:        "",
			Code:          "",
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
		}
	*/
	return c.JSON("product")
}
