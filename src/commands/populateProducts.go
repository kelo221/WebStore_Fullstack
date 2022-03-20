package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"os"
	"path/filepath"
	"strings"
	"webstore/src/database"
	"webstore/src/models"
)

// docker-compose exec backend sh
// go run src/commands/populateProducts.go

func getRandomFruit(fruits []string) string {
	return fruits[int(gofakeit.Price(0, float64(len(fruits))))]
}

func main() {

	var fruits []string

	root := "public/img/products"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".png" {
			fruits = append(fruits, strings.TrimLeft(strings.TrimSuffix(path, filepath.Ext(path)), "public\\img\\products\\"))

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	database.ConnectDB()

	gofakeit.Seed(0)

	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       getRandomFruit(fruits),
			Description: gofakeit.Phrase(),
			Image:       "",
			Price:       gofakeit.Price(0, 100),
		}
		product.Image = "img/products/" + product.Title + ".png"
		database.PushProduct(&product)

	}

}
