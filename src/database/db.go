package database

import (
	"ambassor/src/models"
	"context"
	"flag"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
)

var db driver.Database

var userCol driver.Collection
var productsCol driver.Collection

var orders driver.Collection

func generateCollection(name string) (collection driver.Collection, err error) {

	collection, err = db.Collection(nil, name)
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		collection, err = db.CreateCollection(ctx, name, options)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}

func ConnectDB() {

	var err error
	var client driver.Client
	var conn driver.Connection

	flag.Parse()

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "1234"),
	})

	var dbExists, _ bool

	dbExists, err = client.DatabaseExists(nil, "example")

	if dbExists {
		fmt.Println("That db exists already")

		db, err = client.Database(nil, "example")

		if err != nil {
			log.Fatalf("Failed to open existing database: %v", err)
		}

	} else {
		db, err = client.CreateDatabase(nil, "example", nil)

		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
	}

	userCol, _ = generateCollection("Users")
	productsCol, _ = generateCollection("Products")
	orders, _ = generateCollection("Orders")

}

func AqlNoReturn(query string) {

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)

}

func ReturnArrayOfObject[T any](query string, typeOf []T) (result []T) {

	var dataPayload []T

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc T
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload
	}

	return []T{}

}

func ReturnObject[T any](query string, typeOf T) (result T) {

	var dataPayload []T

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc T
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload[0]
	}

	return T{}

}

func AqlReturnUser(query string) models.User {

	var dataPayload []models.User

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc models.User
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload[0]
	}

	return models.User{}

}

func PushUser(users *models.User) {

	_, err := userCol.CreateDocument(nil, users)

	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	}

}

func PushProduct(product *models.Product) {

	_, err := productsCol.CreateDocument(nil, product)

	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	}

}

func PushShoppingList(list *models.ShoppingCart) {

	_, err := orders.CreateDocument(nil, list)

	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	}

}

/*func AqlReturnProduct(query string) models.Product {

	var dataPayload []models.Product

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc models.Product
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload[0]
	}

	return models.Product{}

}

func AqlReturnOrders(query string) []models.ShoppingCart {

	var dataPayload []models.ShoppingCart

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc models.ShoppingCart
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload
	}

	return []models.ShoppingCart{}

}

func AqlReturnUsers(query string) []models.User {

	var dataPayload []models.User

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc models.User
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		//Do not send hashed password back, less query spaghetti compared to manually setting the query to contain the object
		doc.Password = nil
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload
	}

	return []models.User{}

}*/

/*func AqlReturnProducts(query string) []models.Product {

	var dataPayload []models.Product

	ctx := context.Background()
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		// handle error
	}
	defer func(cursor driver.Cursor) {
		err3 := cursor.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(cursor)
	for {
		var doc models.Product
		_, err2 := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err2) {
			break
		} else if err2 != nil {
			fmt.Println(err2)
		}
		dataPayload = append(dataPayload, doc)
	}

	if len(dataPayload) > 0 {
		return dataPayload
	}

	return []models.Product{}

}*/
