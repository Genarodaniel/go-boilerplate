package repository

import (
	"context"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/repository/user"
)

type UserRepositoryMock struct {
	repository.UserRepository
	SaveUserResponse string
	SaveUserError    error
	GetUserResponse  user.UserDTO
	GetUserError     error
}

func (s UserRepositoryMock) Save(ctx context.Context, user user.UserDTO) (string, error) {
	return s.SaveUserResponse, s.SaveUserError
}

func (s UserRepositoryMock) GetByID(ctx context.Context, userID string) (user.UserDTO, error) {
	return s.GetUserResponse, s.GetUserError
}
