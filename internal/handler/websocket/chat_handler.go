package websocket

import (
	"net/http"

	"balesin-chatEmployee/internal/security"
	"balesin-chatEmployee/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type IncomingMessage struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

func ServeWS(hub *Hub, c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		logger.Log.Error().Msg("Missing token in query parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}

	userID, err := security.ParseToken(token)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to upgrade to websocket")
		return
	}

	// Register user saat connect
	hub.Register(userID, conn)

	defer func() {
		hub.Unregister(userID)
		conn.Close()
		logger.Log.Info().Str("userID", userID).Msg("WebSocket client disconnected")
	}()

	logger.Log.Info().Str("userID", userID).Msg("WebSocket client connected")

	// Read loop - INTI CHAT
	for {
		var msg IncomingMessage
		if err := conn.ReadJSON(&msg); err != nil {
			logger.Log.Warn().Err(err).Str("userID", userID).Msg("WebSocket read error")
			break
		}

		// Lookup target user connection
		targetConn, ok := hub.Get(msg.To)
		if !ok {
			logger.Log.Warn().Str("targetUserID", msg.To).Msg("Target user offline")
			continue
		}

		// Send message to target user
		err := targetConn.WriteJSON(map[string]string{
			"from":    userID,
			"message": msg.Message,
		})

		if err != nil {
			logger.Log.Error().Err(err).Str("targetUserID", msg.To).Msg("WebSocket write error")
		}
	}
}
