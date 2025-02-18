package user

import "context"

type UserServiceMock struct {
	UserServiceInterface
	UserResponse    User
	PostUserError   error
	GetUserResponse GetUserResponse
	GetUserError    error
}

func (s UserServiceMock) PostUser(ctx context.Context, user User) (*User, error) {
	return &s.UserResponse, s.PostUserError
}

func (s UserServiceMock) GetUser(ctx context.Context, userID string) (*User, error) {
	return &s.UserResponse, s.GetUserError
}
