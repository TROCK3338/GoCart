package handlers

import (
	"GoCart/backend/database"
	"GoCart/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCart handles POST /carts to create and add items to a cart
func CreateCart(c *gin.Context) {
	// Get the authenticated user from the context
	user, _ := c.Get("user")
	authenticatedUser := user.(models.User)

	// A temporary struct to bind the request body for adding items to cart
	var requestBody struct {
		ItemIDs []uint `json:"item_ids"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find or create the user's cart
	var cart models.Cart
	result := database.DB.Where("user_id = ?", authenticatedUser.ID).First(&cart)

	if result.Error == gorm.ErrRecordNotFound {
		// If cart doesn't exist, create a new one
		cart = models.Cart{UserID: authenticatedUser.ID, Status: "open"}
		database.DB.Create(&cart)
	}

	// Find the items to add
	var items []models.Item
	database.DB.Find(&items, requestBody.ItemIDs)

	// Add items to the cart
	database.DB.Model(&cart).Association("Items").Append(items)

	c.JSON(http.StatusOK, gin.H{"message": "Items added to cart successfully", "cart": cart})
}

// ListCarts handles GET /carts to list all carts
// NOTE: This endpoint is for demonstration. In a real app, you would
// likely only list the current user's cart or a specific cart by ID.
func ListCarts(c *gin.Context) {
	var carts []models.Cart
	result := database.DB.Preload("Items").Find(&carts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"carts": carts})
}
