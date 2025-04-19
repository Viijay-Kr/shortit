package main

import (
	"fmt"
	"net/http"

	"shortit.sh/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello! from Service Redirect")
	})
	port := fmt.Sprintf(":%s", cfg.ServiceRedirectPort)
	router.Run(port)
}
