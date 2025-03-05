package repository

import (
	"context"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/repository/user"
)

type UserRepositoryMock struct {
	repository.UserRepository
	SaveUserError             error
	GetUserByIDResponse       user.UserDTO
	GetUserByClientIDResponse user.UserDTO
	GetUserByIDError          error
	GetUserClientIDError      error
}

func (s UserRepositoryMock) Save(ctx context.Context, user user.UserDTO) error {
	return s.SaveUserError
}

func (s UserRepositoryMock) GetByID(ctx context.Context, userID string) (user.UserDTO, error) {
	return s.GetUserByIDResponse, s.GetUserByIDError
}

func (s UserRepositoryMock) GetByClientID(ctx context.Context, clientID string) (user.UserDTO, error) {
	return s.GetUserByClientIDResponse, s.GetUserClientIDError
}
