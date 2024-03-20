package main

import (
	"net/http"

	"example.com/agenagn/initializers"
	"example.com/agenagn/middlewares"
	"example.com/agenagn/routes"
	"github.com/gin-gonic/gin"
)

func init() {
   
   initializers.LoadEnvVariables()
   initializers.ConnectToDB()
}


func main() {
	r := gin.Default()
	// initializers.ConnectToDB()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Use(middlewares.Logger())
	r.Use(middlewares.CorsMiddleware())
	
	routes.SetupAuthRoutes(r)
	
	r.Run(":8080")
}