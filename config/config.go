package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Redis struct {
	Host     string
	Port     string
	Password string
}
type Database struct {
	URL      string
	Database string
}
type Config struct {
	Database            Database
	ServiceRedirectPort string
	ServiceGeneratePort string
	ShortitRedirectHost string
	Redis               Redis
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

		go_env := getEnvWithDefault("GO_ENV", "development")

		var shortitRedirectHost string
		shortitRedirectPort := getEnvWithDefault("SERVICE_REDIRECT_PORT", "8002")
		if go_env == "development" {
			shortitRedirectHost = fmt.Sprintf("%s:%s", getEnvWithDefault("SHORTIT_REDIRECT_HOST", "http://localhost"), shortitRedirectPort)
		} else {
			shortitRedirectHost = getEnvWithDefault("SHORTIT_REDIRECT_HOST", "https://shortit.sh")
		}

		config = &Config{
			Database: Database{
				URL:      getEnvWithDefault("DATABASE_URL", "mongodb://root@rootlocalhost:27017/shortit"),
				Database: getEnvWithDefault("DATABASE_NAME", "shortit"),
			},
			ServiceGeneratePort: getEnvWithDefault("SERVICE_GENERATE_PORT", "8001"),
			ServiceRedirectPort: shortitRedirectPort,
			ShortitRedirectHost: shortitRedirectHost,
			Redis: Redis{
				Host:     getEnvWithDefault("REDIS_HOST", "localhost"),
				Port:     getEnvWithDefault("REDIS_PORT", "6379"),
				Password: getEnvWithDefault("REDIS_PASSWORD", ""),
			},
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
