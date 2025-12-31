package chat

import (
	"context"
)

type Service interface {
	SendMessage(ctx context.Context, fromUser, toUser, content string) (*Message, error)
	GetChatHistory(ctx context.Context, chatID string, limit int) ([]*Message, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SendMessage(ctx context.Context, fromUser, toUser, content string) (*Message, error) {
	chat, err := s.repo.GetOrCreateChat(ctx, fromUser, toUser)
	if err != nil {
		return nil, err
	}

	message := &Message{
		ChatID:   chat.ID,
		FromUser: fromUser,
		ToUser:   toUser,
		Content:  content,
	}

	if err := s.repo.SaveMessage(ctx, message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *service) GetChatHistory(ctx context.Context, chatID string, limit int) ([]*Message, error) {
	return s.repo.GetMessages(ctx, chatID, limit)
}
