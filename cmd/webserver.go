package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joncooperworks/jsonjse"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	jse := &jsonjse.JSE{
		Client: &http.Client{},
	}
	mongoDBConnectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	mongoDBClient, err := mongo.NewClient(options.Client().ApplyURI(mongoDBConnectionString))
	if err != nil {
		log.Fatalf("Failed to set up MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoDBClient.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer mongoDBClient.Disconnect(ctx)

	config := &jsonjse.ServerConfig{
		JSE: &jsonjse.JSECache{
			JSE: jse,
			Database: &jsonjse.Database{
				MongoDB: mongoDBClient,
			},
		},
	}
	router := jsonjse.Router(config)
	err = router.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
