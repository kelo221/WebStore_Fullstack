package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"webstore/src/database"
	"webstore/src/models"
)

// docker-compose exec backend sh
// go run src/commands/populateProducts.go
func main() {
	database.ConnectDB()

	gofakeit.Seed(0)

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       gofakeit.Fruit(),
			Description: gofakeit.Phrase(),
			Image:       gofakeit.URL(),
			Price:       gofakeit.Price(0, 100),
		}
		database.PushProduct(&product)
	}

}
