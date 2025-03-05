package mock

import (
	"context"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/app/user"
)

type UserServiceMock struct {
	user.UserServiceInterface
	UserResponse  user.User
	RegisterError error
	GetError      error
	LoginResponse *model.OAuthResponse
	LoginError    error
}

func (s UserServiceMock) Register(ctx context.Context, userRequest model.PostUserRequest) (*user.User, error) {
	return &s.UserResponse, s.RegisterError
}

func (s UserServiceMock) Get(ctx context.Context, userID string) (*user.User, error) {
	return &s.UserResponse, s.GetError
}
