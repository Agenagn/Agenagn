package main

import (
	"net/http"

	"example.com/agenagn/initializers"
	"example.com/agenagn/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initializers.ConnectToDB()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	routes.SetupAuthRoutes(r)

	r.Run(":8080")
}