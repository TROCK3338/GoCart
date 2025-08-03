package handlers

import (
	"GoCart/backend/database"
	"GoCart/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder handles POST /orders to convert a cart to an order
func CreateOrder(c *gin.Context) {
	// Get the authenticated user from the context
	user, _ := c.Get("user")
	authenticatedUser := user.(models.User)

	var requestBody struct {
		CartID uint `json:"cart_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the cart by cart_id and user_id to ensure ownership
	var cart models.Cart
	result := database.DB.Where("id = ? AND user_id = ?", requestBody.CartID, authenticatedUser.ID).First(&cart)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found or does not belong to the user"})
		return
	}

	// Check if the cart is already converted to an order
	if cart.Status == "ordered" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is already converted to an order"})
		return
	}

	// Create a new order
	order := models.Order{
		CartID: cart.ID,
		UserID: authenticatedUser.ID,
	}

	// Begin a transaction for data integrity
	tx := database.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start database transaction"})
		return
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Mark the cart as "ordered" so it cannot be used again
	if err := tx.Model(&cart).Update("status", "ordered").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart status"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// ListOrders handles GET /orders to list all orders for the authenticated user
func ListOrders(c *gin.Context) {
	// Get the authenticated user from the context
	user, _ := c.Get("user")
	authenticatedUser := user.(models.User)

	var orders []models.Order
	result := database.DB.Where("user_id = ?", authenticatedUser.ID).Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
