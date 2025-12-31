package http

import (
	"net/http"

	"balesin-chatEmployee/internal/domain/user"
	"balesin-chatEmployee/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *user.Service
}

func NewAuthHandler(userService *user.Service) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error().Err(err).Msg("Invalid login request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.userService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		logger.Log.Error().Err(err).Str("email", req.Email).Msg("Login failed")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
