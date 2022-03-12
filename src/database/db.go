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

func ConnectDB() {

	var err error
	var client driver.Client
	var conn driver.Connection

	flag.Parse()

	conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://db:8529"},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"),
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

	// Create a collection for users
	userCol, err = db.Collection(nil, "Users")
	if err != nil {
		fmt.Println(err, "creating new...")
		ctx := context.Background()
		options := &driver.CreateCollectionOptions{ /* ... */ }
		userCol, err = db.CreateCollection(ctx, "Users", options)
		if err != nil {
			fmt.Println(err)
		}
	}
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

func AqlJSON(query string) models.User {

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
