package repository

import (
	"context"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/repository/user"
)

type UserRepositoryMock struct {
	repository.UserRepositoryInterface
	SaveUserResponse string
	SaveUserError    error
}

func (s UserRepositoryMock) Save(ctx context.Context, user user.User) (string, error) {
	return s.SaveUserResponse, s.SaveUserError
}
