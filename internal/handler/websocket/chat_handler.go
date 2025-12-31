package websocket

import (
	"net/http"

	"balesin-chatEmployee/internal/security"
	"balesin-chatEmployee/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // nanti bisa dipersempit
	},
}

type ChatHandler struct {
	hub *Hub
}

func NewChatHandler(hub *Hub) *ChatHandler {
	return &ChatHandler{hub: hub}
}

func (h *ChatHandler) HandleChat(c *gin.Context) {
	// 👉 Get token from query parameter
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

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// REGISTER USER
	h.hub.Register(userID, conn)
	defer func() {
		h.hub.Unregister(userID)
		conn.Close()
		logger.Log.Info().Str("userID", userID).Msg("WebSocket client disconnected")
	}()

	logger.Log.Info().Str("userID", userID).Msg("WebSocket client connected")

	for {
		var msg IncomingMessage
		if err := conn.ReadJSON(&msg); err != nil {
			logger.Log.Warn().Err(err).Str("userID", userID).Msg("WebSocket read error")
			break
		}

		// CARI USER TUJUAN
		targetConn, ok := h.hub.Get(msg.To)
		if !ok {
			logger.Log.Warn().Str("targetUserID", msg.To).Msg("Target user offline")
			continue
		}

		out := OutgoingMessage{
			From:    userID,
			Message: msg.Message,
		}

		// KIRIM KE USER TUJUAN
		if err := targetConn.WriteJSON(out); err != nil {
			logger.Log.Error().Err(err).Str("targetUserID", msg.To).Msg("WebSocket write error")
		}
	}
}

