package user

import "context"

type UserServiceSpy struct {
	UserServiceInterface
	PostUserResponse PostUserResponse
	PostUserError    error
}

func (s UserServiceSpy) PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error) {
	return &s.PostUserResponse, s.PostUserError
}
