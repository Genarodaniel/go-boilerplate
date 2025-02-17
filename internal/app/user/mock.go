package user

import "context"

type UserServiceMock struct {
	UserServiceInterface
	PostUserResponse PostUserResponse
	PostUserError    error
}

func (s UserServiceMock) PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error) {
	return &s.PostUserResponse, s.PostUserError
}
