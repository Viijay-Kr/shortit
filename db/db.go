package db

import (
	"context"
	"fmt"

	"github.com/Viijay-Kr/shortit/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Initialize() error {
	cfg := config.GetConfig()

	var uri string
	if uri = cfg.DatabaseURL; uri == "" {
		return fmt.Errorf("database URL is not set")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	mongo_client, err := mongo.Connect(clientOptions)

	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	client = mongo_client

	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return nil
}

func GetClient() *mongo.Client {
	if client == nil {
		panic("MongoDB client is not initialized. Call Initialize() first.")
	}
	return client
}
