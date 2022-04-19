package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"strconv"
	"webstore/src/database"
	"webstore/src/models"
)

// docker-compose exec backend sh
// go run src/commands/populateOrders.go
func main() {
	database.ConnectDB()

	gofakeit.Seed(0)

	for i := 0; i < 10; i++ {

		var products []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {

			products = append(products, models.OrderItem{
				ProductTitle: gofakeit.Fruit(),
				Price:        gofakeit.Price(0, 50),
				Quantity:     uint(rand.Intn(5)),
			})
		}

		orderComplete := models.ShoppingCart{
			UserId:     strconv.Itoa((rand.Intn(10)) + 1),
			FirstName:  gofakeit.Name(),
			LastName:   gofakeit.LastName(),
			Email:      gofakeit.Email(),
			Address:    gofakeit.Address().Street,
			City:       gofakeit.City(),
			Country:    gofakeit.Country(),
			Zip:        gofakeit.Zip(),
			Complete:   true,
			OrderItems: products,
		}

		fmt.Printf("%v", orderComplete)
		database.PushShoppingList(&orderComplete)
	}

}
