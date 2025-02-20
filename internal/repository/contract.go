package repository

import (
	"context"
	"go-boilerplate/internal/repository/user"
)

type UserRepository interface {
	Save(ctx context.Context, order user.UserDTO) (string, error)
	GetByID(ctx context.Context, userID string) (user.UserDTO, error)
	// Update(ctx context.Context, order user.UserDTO) error
	// Delete(ctx context.Context, userID string) error
}
