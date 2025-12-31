package main

import (
	"balesin-chatEmployee/internal/database"
	"balesin-chatEmployee/internal/handler/http"
	"balesin-chatEmployee/internal/handler/websocket"
	"balesin-chatEmployee/internal/middleware"
	"balesin-chatEmployee/internal/repository"
	"balesin-chatEmployee/internal/service"
	"balesin-chatEmployee/pkg/logger"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Init("balesin-chatEmployee")

	logger.Log.Info().Msg("Starting server...")
	r := gin.Default()

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())

	r.Use(logger.HTTPLogger())

	database.Connect()
	logger.Log.Info().Msg("Connected to database")

	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := http.NewAuthHandler(authService)

	hub := websocket.NewHub()

	auth.GET("/me", func(c *gin.Context) {
		userID := c.GetString("user_id")
		c.JSON(200, gin.H{
			"user_id": userID,
		})
	})

	r.POST("/auth/login", authHandler.Login)
	r.GET("/ws/chat", func(c *gin.Context) {
		websocket.ServeWS(hub, c)
	})

	logger.Log.Info().Msg("========================================")
	logger.Log.Info().Msg("Server is running successfully!")
	logger.Log.Info().Str("port", "8080").Msg("Listening on http://localhost:8080")
	logger.Log.Info().Msg("Ready to accept requests")
	logger.Log.Info().Msg("========================================")
	r.Run(":8080")
}
