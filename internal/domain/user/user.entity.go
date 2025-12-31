package user

import (
	"time"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
}
