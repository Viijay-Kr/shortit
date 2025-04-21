package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viijay-Kr/shortit/cache"
	"github.com/Viijay-Kr/shortit/config"
	"github.com/Viijay-Kr/shortit/core"
	"github.com/Viijay-Kr/shortit/db"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()

	redis_err := cache.Initialize()
	db_err := db.Initialize()

	if db_err != nil {
		panic(fmt.Sprintf("Failed to initialize DB: %v", db_err))
	}
	if redis_err != nil {
		panic(fmt.Sprintf("Failed to initialize Redis: %v", redis_err))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, from service generate!")
	})

	router.POST("/api/generate", generateShortURL)

	port := fmt.Sprintf(":%s", cfg.ServiceGeneratePort)

	router.Run(port)
}

// Implementation of generateShortURL function which is post handler for gin route
// This is a post handler that receives a big url in request body and returns a short url
func generateShortURL(c *gin.Context) {
	// Get the big url from request body
	big_url := c.PostForm("url")

	short_url, err := core.GenerateShortUrl(big_url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cache.Set(short_url.ID, short_url.Sanitized)

	c.JSON(http.StatusOK, gin.H{"short_url": short_url.Shortened})
}
