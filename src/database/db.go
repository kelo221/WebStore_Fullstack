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

// **************************************************

func PrintCollection(db driver.Database, name string) {

	var err error
	var cursor driver.Cursor

	querystring := "FOR doc IN users LIMIT 10 RETURN doc"

	cursor, err = db.Query(nil, querystring, nil)

	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	defer cursor.Close()

	for {
		var doc models.User
		var metadata driver.DocumentMeta

		metadata, err = cursor.ReadDocument(nil, &doc)

		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Fatalf("Doc returned: %v", err)
		} else {
			fmt.Print("Dot doc ", metadata, doc, "\n")
		}
	}
}

func PushUser(users *models.User) {

	_, err := userCol.CreateDocument(nil, users)

	if err != nil {
		log.Fatalf("Failed to create documents: %v", err)
	}

}
