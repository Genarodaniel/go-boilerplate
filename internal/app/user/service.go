package user

import (
	"context"
	"database/sql"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/repository"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/internal/services/kafka"
	"go-boilerplate/pkg/cryptography"
	"go-boilerplate/pkg/customerror"
)

type UserServiceInterface interface {
	Register(ctx context.Context, userRequest model.PostUserRequest) (*User, error)
	Get(ctx context.Context, userID string) (*User, error)
}

type UserService struct {
	KafkaProducer  kafka.KafkaInterface
	UserRepository repository.UserRepository
}

func NewUserService(kafkaProducer kafka.KafkaInterface, userRepository repository.UserRepository) *UserService {
	return &UserService{
		KafkaProducer:  kafkaProducer,
		UserRepository: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, userRequest model.PostUserRequest) (*User, error) {
	user, err := NewUser(
		userRequest.Email,
		userRequest.Name,
	)
	if err != nil {
		return nil, customerror.NewValidationError(err.Error())
	}

	secretHashed, err := cryptography.HashSecret(user.ClientSecret)
	if err != nil {
		return nil, customerror.NewApplicationError(err.Error())
	}

	userDto := userRepository.UserDTO{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		ClientSecret: secretHashed,
		ClientID:     user.ClientID,
	}

	if err := s.UserRepository.Save(ctx, userDto); err != nil {
		return nil, customerror.NewApplicationError(err.Error())
	}

	userQueueRequest := model.UserQueue{
		Name:     user.Name,
		ClientID: user.ClientID,
		Email:    user.Email,
	}

	if err := s.KafkaProducer.Produce(ctx, "users", "user.create", userQueueRequest); err != nil {
		return nil, customerror.NewApplicationError(err.Error())
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, userID string) (*User, error) {
	userDto, err := s.UserRepository.GetByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerror.NewNotFoundError("user not found")
		}

		return nil, customerror.NewApplicationError(err.Error())
	}

	return &User{
		ID:    userDto.ID,
		Email: userDto.Email,
		Name:  userDto.Name,
	}, nil

}
