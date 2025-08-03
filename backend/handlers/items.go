package handlers

import (
	"GoCart/backend/database"
	"GoCart/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateItem handles POST /items to create a new item
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&item)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully", "item": item})
}

// ListItems handles GET /items to list all available items
func ListItems(c *gin.Context) {
	var items []models.Item
	result := database.DB.Find(&items)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}
