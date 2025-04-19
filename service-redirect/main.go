package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Viijay-Kr/shortit/cache"
	"github.com/Viijay-Kr/shortit/config"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()
	redis_err := cache.Initialize()

	if redis_err != nil {
		fmt.Println("Error initializing Redis:", redis_err)
		panic(redis_err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello! from Service Redirect")
	})
	router.GET("/:id", getShortUrl)

	port := fmt.Sprintf(":%s", cfg.ServiceRedirectPort)
	router.Run(port)
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
		c.String(http.StatusNotFound, "URL not found")
		return
	}

	c.Redirect(http.StatusFound, val)
}
