package main

import (
	"time"

	"drive-mini/internal/database"
	"drive-mini/internal/handlers"
	"drive-mini/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	handlers.SeedUsers()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/me", handlers.Me)
		auth.GET("/files", handlers.GetFiles)
		auth.POST("/files", handlers.CreateFile)
		auth.PUT("/files/:id", handlers.UpdateFile)
		auth.DELETE("/files/:id", handlers.DeleteFile)
	}

	r.Run(":8080")
}
