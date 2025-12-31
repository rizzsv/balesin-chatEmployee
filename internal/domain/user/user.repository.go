package user

import (
	"context"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
}
