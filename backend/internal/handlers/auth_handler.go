package handlers

import (
	"net/http"

	"drive-mini/internal/auth"
	"drive-mini/internal/database"
	"drive-mini/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	createIfNotExist("admin@drive.com", "admin123", models.RoleAdmin)
	createIfNotExist("user1@drive.com", "user123", models.RoleUser)
	createIfNotExist("user2@drive.com", "user123", models.RoleUser)
}

func createIfNotExist(email, pass string, role models.Role) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)

	database.DB.Create(&models.User{
		ID:       email,
		Email:    email,
		Password: string(hash),
		Role:     role,
	})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string
		Password string
	}
	c.BindJSON(&req)

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := auth.GenerateToken(user.ID, string(user.Role))
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetString("user_id"),
		"role":    c.GetString("role"),
	})
}
