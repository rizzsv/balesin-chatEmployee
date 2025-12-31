package websocket

import (
	"net/http"

	"balesin-chatEmployee/internal/security"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware validates JWT token from query parameter for WebSocket
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing token"})
			return
		}

		userID, err := security.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
