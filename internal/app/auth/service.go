package auth

import (
	"context"
	"fmt"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/repository"
	"go-boilerplate/pkg/cryptography"
	"go-boilerplate/pkg/customerror"
)

type AuthServiceInterface interface {
	Authenticate(ctx context.Context, clientID, clientSecret, grantType string) (Auth, error)
}

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (a *AuthService) Authenticate(ctx context.Context, clientID, clientSecret, grantType string) (Auth, error) {
	user, err := a.userRepository.GetByClientID(ctx, clientID)
	if err != nil {
		return Auth{}, customerror.NewUnauthorizedError(model.ErrInvalidCredentials.Error())
	}

	if !cryptography.ValidateClientSecret(clientSecret, user.ClientSecret) {
		return Auth{}, customerror.NewUnauthorizedError(model.ErrInvalidCredentials.Error())
	}

	// token, err := a.Jwt.GenerateToken(clientID)
	// if err != nil {
	// 	return Auth{}, err
	// }

	fmt.Println(user)

	return Auth{
		AccessToken: "12",
	}, nil
}
