package main

import (
	"log"

	"balesin-chatEmployee/internal/config"
	"balesin-chatEmployee/internal/domain/user"
	"balesin-chatEmployee/internal/middleware"
	"balesin-chatEmployee/internal/repository/postgres"
	"balesin-chatEmployee/internal/shared"
	"balesin-chatEmployee/internal/transport/http"
	"balesin-chatEmployee/internal/transport/websocket"
	"balesin-chatEmployee/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize logger
	logger.Init("balesin-chatEmployee")
	logger.Log.Info().Msg("Starting server...")

	// Load configurations
	config.LoadJWT()
	config.ConnectDatabase()

	// Initialize Gin
	r := gin.Default()
	r.Use(logger.HTTPLogger())

	// Initialize repositories
	userRepo := postgres.NewUserRepository()

	// Initialize services
	userService := user.NewService(userRepo)

	// Initialize handlers
	authHandler := http.NewAuthHandler(userService)

	// WebSocket hub
	hub := websocket.NewHub()
	chatHandler := websocket.NewChatHandler(hub)

	// Public routes
	r.POST("/auth/login", authHandler.Login)
	r.GET("/ws/chat", chatHandler.HandleChat)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/me", func(c *gin.Context) {
			userID := c.GetString(shared.ContextKeyUserID)
			c.JSON(200, gin.H{
				"user_id": userID,
			})
		})
	}

	// Start server
	logger.Log.Info().Msg("========================================")
	logger.Log.Info().Msg("Server is running successfully!")
	logger.Log.Info().Str("port", "8080").Msg("Listening on http://localhost:8080")
	logger.Log.Info().Msg("Ready to accept requests")
	logger.Log.Info().Msg("========================================")

	if err := r.Run(":8080"); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
}
