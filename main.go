package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-gin-start/internal/auth"
	"go-gin-start/internal/db"
	"go-gin-start/internal/handler"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if databaseURL == "" || jwtSecret == "" {
		log.Fatal("DATABASE_URL and JWT_SECRET must be set")
	}

	// Connect to DB
	err := db.InitDB(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.DB.Close()

	// Run migrations
	// if err := db.AutoMigrate(databaseURL); err != nil {
	// 	log.Fatalf("Migration failed: %v", err)
	// }

	// Setup Gin
	r := gin.Default()

	r.GET("/hello", handler.HelloHandler())

	// Auth routes
	r.POST("/login", handler.LoginHandler(jwtSecret))

	// Protected routes example
	protected := r.Group("/api")
	protected.Use(auth.AuthMiddleware(jwtSecret))
	protected.GET("/me", handler.MeHandler())

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
