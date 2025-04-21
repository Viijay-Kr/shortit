package db

import (
	"context"
	"fmt"

	"github.com/Viijay-Kr/shortit/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func Initialize() error {
	cfg := config.GetConfig()

	var uri string
	if uri = cfg.DatabaseURL; uri == "" {
		return fmt.Errorf("Database URL is not set")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return fmt.Errorf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
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
