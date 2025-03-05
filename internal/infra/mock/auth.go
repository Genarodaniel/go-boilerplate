package mock

import (
	"context"
	"go-boilerplate/internal/app/auth"
)

type AuthServiceMock struct {
	auth.AuthServiceInterface
	AuthenticateResponse auth.Auth
	AuthenticateError    error
}

func (s AuthServiceMock) Authenticate(ctx context.Context, clientID, clientSecret, grantType string) (auth.Auth, error) {
	return s.AuthenticateResponse, s.AuthenticateError
}
