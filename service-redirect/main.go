package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viijay-Kr/shortit/cache"
	"github.com/Viijay-Kr/shortit/config"
	"github.com/Viijay-Kr/shortit/db"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()
	redis_err := cache.Initialize()
	if redis_err != nil {
		fmt.Println("Error initializing Redis:", redis_err)
		panic(redis_err)
	}

	db_err := db.Initialize()
	if db_err != nil {
		fmt.Println("Error initializing DB:", db_err)
		panic(db_err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello! from Service Redirect")
	})
	router.GET("/:id", getShortUrl)

	port := fmt.Sprintf(":%s", cfg.ServiceRedirectPort)
	router.Run(port)
}

type ShortUrl struct {
	Hash      string `bson:"hash,omitempty"`
	Sanitized string `bson:"sanitized,omitempty"`
}

// Implementation of getShortUrl function
func getShortUrl(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	val, err := cache.Get(id)

	if err != nil {
		fmt.Println("Error getting value from cache:", err)
	}

	if val != "" {
		fmt.Println("Value found in cache")
		c.Redirect(http.StatusFound, val)
		return
	}

	// If not found in cache, check the database
	result, err := db.GetLongUrl(id)
	if err != nil {
		fmt.Println("Error getting value from DB:", err)
		c.String(http.StatusNotFound, "URL not found")
		return
	}
	// Set the value in cache
	err = cache.Set(id, result)
	if err != nil {
		fmt.Println("Error setting value in cache:", err)
	}
	c.Redirect(http.StatusFound, result)
}
