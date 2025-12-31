package config

import (
	"os"
)

var JWTSecret string

func LoadJWT() {
	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		JWTSecret = "SUPER_SECRET_KEY" // fallback
	}
}

func GetJWTSecret() string {
	return JWTSecret
}
