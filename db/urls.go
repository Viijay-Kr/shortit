package db

import (
	"context"
	"fmt"

	"github.com/Viijay-Kr/shortit/config"
	"github.com/Viijay-Kr/shortit/core"
	"go.mongodb.org/mongo-driver/bson"
)

type ShortUrl struct {
	Hash      string `bson:"hash,omitempty"`
	Sanitized string `bson:"sanitized,omitempty"`
	Shortened string `bson:"shortened,omitempty"`
}

func GetLongUrl(hash string) (string, error) {
	cfg := config.GetConfig()
	client := GetClient()
	collection := client.Database(cfg.Database.Database).Collection("urls")
	filter := bson.M{"hash": hash}

	var shortUrl ShortUrl
	res := collection.FindOne(context.TODO(), filter)

	err := res.Decode(&shortUrl)

	if err != nil {
		fmt.Println("Error decoding result:", err)
		return "", err
	}

	return shortUrl.Sanitized, nil
}

func InsertUrl(shortUrl core.ShortUrl) error {
	cfg := config.GetConfig()
	client := GetClient()
	collection := client.Database(cfg.Database.Database).Collection("urls")

	_, err := collection.InsertOne(context.TODO(), shortUrl)
	if err != nil {
		fmt.Println("Error inserting URL:", err)
		return err
	}

	return nil
}
