package websocket

import (
	"net/http"

	"balesin-chatEmployee/internal/config"
	"balesin-chatEmployee/internal/security"
	"balesin-chatEmployee/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	hub *Hub
}

func NewChatHandler(hub *Hub) *ChatHandler {
	return &ChatHandler{hub: hub}
}

type IncomingMessage struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type OutgoingMessage struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func (h *ChatHandler) HandleChat(c *gin.Context) {
	// Get token from query parameter
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}

	// Validate JWT token
	userID, err := security.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Upgrade to WebSocket
	upgrader := config.GetWSUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to upgrade connection")
		return
	}

	// Register user connection
	h.hub.Register(userID, conn)
	defer func() {
		h.hub.Unregister(userID)
		conn.Close()
		logger.Log.Info().Str("userID", userID).Msg("WebSocket client disconnected")
	}()

	logger.Log.Info().Str("userID", userID).Msg("WebSocket client connected")

	// Read messages from client
	for {
		var msg IncomingMessage
		if err := conn.ReadJSON(&msg); err != nil {
			logger.Log.Warn().Err(err).Str("userID", userID).Msg("WebSocket read error")
			break
		}

		// Find target user connection
		targetConn, ok := h.hub.Get(msg.To)
		if !ok {
			logger.Log.Warn().Str("targetUserID", msg.To).Msg("Target user offline")
			continue
		}

		// Send message to target user
		out := OutgoingMessage{
			From:    userID,
			Message: msg.Message,
		}

		if err := targetConn.WriteJSON(out); err != nil {
			logger.Log.Error().Err(err).Str("targetUserID", msg.To).Msg("WebSocket write error")
		}
	}
}
