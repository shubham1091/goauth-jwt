package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shubham1091/goauth-jwt/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "pong"})
	})


	routes.AuthRoutes(router)
	routes.UserRoutes(router)


	router.Run(":" + port)
}
