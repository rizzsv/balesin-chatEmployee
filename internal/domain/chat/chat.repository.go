package chat

import (
	"context"
)

type Repository interface {
	SaveMessage(ctx context.Context, message *Message) error
	GetMessages(ctx context.Context, chatID string, limit int) ([]*Message, error)
	GetOrCreateChat(ctx context.Context, user1, user2 string) (*Chat, error)
}
