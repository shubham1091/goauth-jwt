package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/shubham1091/goauth-jwt/controllers"
	"github.com/shubham1091/goauth-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}