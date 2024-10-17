package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	helpers "github.com/shubham1091/goauth-jwt/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		autharizaton := c.Request.Header.Get("Authorization")

		if autharizaton == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization header provided"})
			c.Abort()
		}

		clientToken := strings.Split(autharizaton, "Bearer ")[1]

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("user_type", claims.User_type)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
