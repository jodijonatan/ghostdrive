package handlers

import (
	"net/http"
	"time"

	"drive-mini/internal/database"
	"drive-mini/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetFiles(c *gin.Context) {
	role := c.GetString("role")
	userID := c.GetString("user_id")

	files := []models.FileMeta{}

	if role == "ADMIN" {
		database.DB.Find(&files)
	} else {
		database.DB.Where("owner_id = ?", userID).Find(&files)
	}

	c.JSON(http.StatusOK, files)
}

func CreateFile(c *gin.Context) {
	userID := c.GetString("user_id")

	var req struct {
		Filename string `json:"filename"`
	}
	c.BindJSON(&req)

	file := models.FileMeta{
		ID:        uuid.NewString(),
		Filename:  req.Filename,
		OwnerID:   userID,
		CreatedAt: time.Now(),
	}

	database.DB.Create(&file)
	c.JSON(http.StatusOK, file)
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")
	role := c.GetString("role")

	var file models.FileMeta
	if err := database.DB.Where("id = ?", id).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if role != "ADMIN" && file.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	var req struct {
		Filename string `json:"filename"`
	}
	c.BindJSON(&req)

	file.Filename = req.Filename
	database.DB.Save(&file)

	c.JSON(http.StatusOK, file)
}

func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")
	role := c.GetString("role")

	var file models.FileMeta
	if err := database.DB.Where("id = ?", id).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if role != "ADMIN" && file.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	database.DB.Delete(&file)
	c.Status(http.StatusOK)
}
