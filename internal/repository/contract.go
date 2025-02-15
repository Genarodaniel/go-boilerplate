package repository

import (
	"context"
	"go-boilerplate/internal/repository/user"
)

type UserRepositoryInterface interface {
	Save(ctx context.Context, order user.User) (string, error)
}
