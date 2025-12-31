package postgres

import (
	"context"
	"errors"

	"balesin-chatEmployee/internal/config"
	"balesin-chatEmployee/internal/domain/user"
)

type userRepository struct{}

func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `
		SELECT id, email, password_hash, role, is_active, created_at
		FROM users
		WHERE email = $1
	`
	row := config.DB.QueryRow(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.IsActive, &u.CreatedAt)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &u, nil
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, role, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := config.DB.Exec(ctx, query, u.ID, u.Email, u.PasswordHash, u.Role, u.IsActive, u.CreatedAt)
	return err
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	query := `
		SELECT id, email, password_hash, role, is_active, created_at
		FROM users
		WHERE id = $1
	`
	row := config.DB.QueryRow(ctx, query, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.IsActive, &u.CreatedAt)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &u, nil
}
