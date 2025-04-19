package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shortit.sh/config"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, from service generate!")
	})

	port := fmt.Sprintf(":%s", cfg.ServiceGeneratePort)
	router.Run(port)
}
