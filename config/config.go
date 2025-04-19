package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL         string
	RedisURL            string
	ServiceRedirectPort string
	ServiceGeneratePort string
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found: %v", err)
		}

		config = &Config{
			DatabaseURL:         getEnvWithDefault("DATABASE_URL", "mongodb://root@rootlocalhost:27017/shortit"),
			RedisURL:            getEnvWithDefault("REDIS_URL", "redis://localhost:6379"),
			ServiceRedirectPort: getEnvWithDefault("SERVICE_REDIRECT_PORT", "8002"),
			ServiceGeneratePort: getEnvWithDefault("SERVICE_GENERATE_PORT", "8001"),
		}
	})

	return config
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
