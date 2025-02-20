package mocks

import (
	"context"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/app/user"
)

type UserServiceMock struct {
	user.UserServiceInterface
	UserResponse    user.User
	PostUserError   error
	GetUserResponse model.GetUserResponse
	GetUserError    error
}

func (s UserServiceMock) PostUser(ctx context.Context, user user.User) (*user.User, error) {
	return &s.UserResponse, s.PostUserError
}

func (s UserServiceMock) GetUser(ctx context.Context, userID string) (*user.User, error) {
	return &s.UserResponse, s.GetUserError
}
