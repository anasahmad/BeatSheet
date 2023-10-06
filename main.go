package main

import (
	_ "BeatSheet/docs"
	"BeatSheet/internal/server"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//	@title			BeatSheet Swagger
//	@version		1.0
//	@description	This is a beat sheet service

//	@contact.name	Anas Ahmad
//	@contact.url	https://github.com/anasahmad
//	@contact.email	anasahmaddev@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v0/beatsheet/

// @securityDefinitions.basic	BasicAuth
func main() {
	dbInit()
	server.Serve()
}

func dbInit() {
	// Define MongoDB connection options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	ctx := context.TODO()

	// Create a MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new database
	db := client.Database("beatsheetdb")

	// Create collection (if it doesn't exist)
	err = db.CreateCollection(ctx, "beatsheets")
	if err != nil {
		log.Println("BeatSheets collection already exists")
	}

	//// Close the MongoDB client when done
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}
