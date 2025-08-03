package middleware

import (
	"net/http"
	"GoCart/backend/database"
	"GoCart/backend/models"

	"github.com/gin-gonic/gin"
)

// AuthRequired is a middleware to check for a valid user token
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		var user models.User
		// Find the user by the token
		result := database.DB.Where("token = ?", token).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store the user in the context for later use in handlers
		c.Set("user", user)
		
		c.Next()
	}
}