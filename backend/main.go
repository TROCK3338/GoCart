package main

import (
	"GoCart/backend/database"
	"GoCart/backend/handlers"
	"GoCart/backend/middleware"
	"log"

	"github.com/gin-contrib/cors" // Import the cors package
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// Configure and use CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true           // In a real application, you'd be more restrictive, e.g., config.AllowOrigins = []string{"http://localhost:3000"}
	config.AddAllowHeaders("Authorization") // This is crucial for passing the token
	router.Use(cors.New(config))

	// Public endpoints
	router.POST("/users", handlers.CreateUser)
	router.POST("/users/login", handlers.Login)
	router.GET("/items", handlers.ListItems)
	router.POST("/items", handlers.CreateItem)

	// Protected endpoints
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/carts", handlers.CreateCart)
		protected.GET("/carts", handlers.ListCarts)
		protected.POST("/orders", handlers.CreateOrder)
		protected.GET("/orders", handlers.ListOrders)
	}

	log.Println("Server starting on port 8080...")
	router.Run(":8080")
}
