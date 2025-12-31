package postgres

import (
	"context"
	"time"

	"balesin-chatEmployee/internal/config"
	"balesin-chatEmployee/internal/domain/chat"

	"github.com/google/uuid"
)

type chatRepository struct{}

func NewChatRepository() chat.Repository {
	return &chatRepository{}
}

func (r *chatRepository) SaveMessage(ctx context.Context, message *chat.Message) error {
	if message.ID == "" {
		message.ID = uuid.New().String()
	}
	if message.CreatedAt.IsZero() {
		message.CreatedAt = time.Now()
	}

	query := `
		INSERT INTO messages (id, chat_id, from_user, to_user, content, created_at, is_read)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := config.DB.Exec(ctx, query,
		message.ID, message.ChatID, message.FromUser,
		message.ToUser, message.Content, message.CreatedAt, message.IsRead,
	)
	return err
}

func (r *chatRepository) GetMessages(ctx context.Context, chatID string, limit int) ([]*chat.Message, error) {
	query := `
		SELECT id, chat_id, from_user, to_user, content, created_at, is_read
		FROM messages
		WHERE chat_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`
	rows, err := config.DB.Query(ctx, query, chatID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*chat.Message
	for rows.Next() {
		var msg chat.Message
		err := rows.Scan(&msg.ID, &msg.ChatID, &msg.FromUser, &msg.ToUser, &msg.Content, &msg.CreatedAt, &msg.IsRead)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	return messages, nil
}

func (r *chatRepository) GetOrCreateChat(ctx context.Context, user1, user2 string) (*chat.Chat, error) {
	// Check if chat already exists
	query := `
		SELECT id, participant1, participant2, created_at
		FROM chats
		WHERE (participant1 = $1 AND participant2 = $2)
		   OR (participant1 = $2 AND participant2 = $1)
		LIMIT 1
	`
	var c chat.Chat
	err := config.DB.QueryRow(ctx, query, user1, user2).Scan(&c.ID, &c.Participant1, &c.Participant2, &c.CreatedAt)
	if err == nil {
		return &c, nil
	}

	// Create new chat
	c.ID = uuid.New().String()
	c.Participant1 = user1
	c.Participant2 = user2
	c.CreatedAt = time.Now()

	insertQuery := `
		INSERT INTO chats (id, participant1, participant2, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err = config.DB.Exec(ctx, insertQuery, c.ID, c.Participant1, c.Participant2, c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
