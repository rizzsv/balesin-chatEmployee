package main

import (
	"context"
	"log"

	"balesin-chatEmployee/internal/database"
	"balesin-chatEmployee/internal/security"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	hash, _ := security.HashPassword("password123")

	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, 'admin')
	`, "admin@company.com", hash)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("seed user created")
}
