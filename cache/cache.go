package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/Viijay-Kr/shortit/config"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Initialize() error {
	cfg := config.GetConfig()

	redis_addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	redis_pass := cfg.Redis.Password
	client = redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	if val != "" {
		fmt.Println("Value found in cache")
	}
	return val, nil
}

func Set(key, value string) error {
	ctx := context.Background()
	err := client.Set(ctx, key, value, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Failed to set value in cache:", err)
		return err
	}
	fmt.Println("Value set in cache for key and value", key, value)
	return nil
}
