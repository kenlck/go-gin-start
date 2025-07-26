// Package handler provides authentication endpoints.
package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go-gin-start/internal/auth"
	"go-gin-start/internal/db"
)

// LoginHandler handles user login and JWT issuance.
func LoginHandler(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		repo := db.UserRepo{}
		user, err := repo.FindByUsername(context.Background(), req.Username)
		if err != nil || !auth.CheckPassword(user.PasswordHash, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		token, err := auth.GenerateJWT(user.ID, jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// MeHandler returns the current authenticated user.
func MeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("user_id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		repo := db.UserRepo{}
		user, err := repo.FindByUsername(context.Background(), os.Getenv("username"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":       user.ID,
			"username": user.Username,
		})
	}
}
