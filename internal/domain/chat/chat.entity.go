package chat

import (
	"time"
)

type Chat struct {
	ID           string
	Participant1 string // userID
	Participant2 string // userID
	CreatedAt    time.Time
}
