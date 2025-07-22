package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/davidperjans/filmradar/internal/auth"
	"github.com/davidperjans/filmradar/internal/db"
)

func main() {
	// Ladda miljövariabler
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Anslut till databasen
	db.Connect()

	// Skapa router
	router := gin.Default()

	// Auth routes
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	// Skyddad route (kräver JWT)
	protected := router.Group("/api")
	protected.Use(auth.RequireAuth()) // 👈 JWT middleware
	protected.GET("/me", func(c *gin.Context) {
		userID := c.MustGet("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	// Test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Kör servern
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
