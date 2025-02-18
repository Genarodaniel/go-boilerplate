package repository

import (
	"context"
	"go-boilerplate/internal/repository/user"
)

type UserRepositoryInterface interface {
	Save(ctx context.Context, order user.UserDTO) (string, error)
	GetByID(ctx context.Context, userID string) (user.UserDTO, error)
}
