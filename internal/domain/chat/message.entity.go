package chat

import (
	"time"
)

type Message struct {
	ID        string
	ChatID    string
	FromUser  string
	ToUser    string
	Content   string
	CreatedAt time.Time
	IsRead    bool
}
