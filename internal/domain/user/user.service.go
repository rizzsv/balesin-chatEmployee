package user

import (
	"context"
	"errors"

	"balesin-chatEmployee/internal/security"
	"balesin-chatEmployee/pkg/logger"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		logger.Log.Error().Err(err).Str("email", email).Msg("User not found")
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", errors.New("user is inactive/disabled")
	}

	if !security.CheckPassword(password, user.PasswordHash) {
		logger.Log.Info().Str("email", email).Msg("Invalid password attempt")
		return "", errors.New("invalid credentials")
	}

	token, err := security.GeneratorToken(user.ID)
	if err != nil {
		logger.Log.Error().Err(err).Str("userID", user.ID).Msg("Failed to generate token")
		return "", errors.New("failed to generate token")
	}

	logger.Log.Info().Str("userID", user.ID).Msg("User logged in successfully")
	return token, nil
}
