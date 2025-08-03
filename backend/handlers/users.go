package handlers

import (
	"GoCart/backend/database"
	"GoCart/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// A temporary struct for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var requestBody CreateUserRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Clear the password before sending the response
	user.Password = ""

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// Login for existing user based on username and password
func Login(c *gin.Context) {
	var loginUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := database.DB.Where("username = ?", loginUser.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.Password != loginUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := "some-generated-token-for-" + loginUser.Username
	user.Token = token
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
