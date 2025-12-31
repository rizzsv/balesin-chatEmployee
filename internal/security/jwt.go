package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("7f8c2a9e4d6b1f3a0c9e8b5d2a4f6c1e9b7a3d5f8e2c4a6b0d1f9e7a3c5")

func GeneratorToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
