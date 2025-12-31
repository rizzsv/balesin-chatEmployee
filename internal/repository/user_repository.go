package repository

import (
	"context"
	"errors"

	"balesin-chatEmployee/internal/database"
	"balesin-chatEmployee/internal/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, role, is_active, created_at
		FROM users
		WHERE email = $1
	`
	row := database.DB.QueryRow(ctx, query, email)

	var user domain.User

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.IsActive, &user.CreatedAt)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
