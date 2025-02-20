package user

import (
	"context"
	"database/sql"
	"go-boilerplate/internal/infra/errhandler"
	"go-boilerplate/internal/repository"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/internal/services/kafka"
)

type UserServiceInterface interface {
	PostUser(ctx context.Context, user User) (*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
}

type UserService struct {
	KafkaProducer  kafka.KafkaInterface
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(kafkaProducer kafka.KafkaInterface, userRepository repository.UserRepositoryInterface) *UserService {
	return &UserService{
		KafkaProducer:  kafkaProducer,
		UserRepository: userRepository,
	}
}

func (s *UserService) PostUser(ctx context.Context, user User) (*User, error) {
	if err := user.Validate(); err != nil {
		return nil, errhandler.NewValidationError(err.Error())
	}

	userDto := userRepository.UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}

	userID, err := s.UserRepository.Save(ctx, userDto)
	if err != nil {
		return nil, errhandler.NewApplicationError(err.Error())
	}

	user.ID = userID
	if err := s.KafkaProducer.Produce(ctx, "users", "user.create", user); err != nil {
		return nil, errhandler.NewApplicationError(err.Error())
	}

	return &user, nil
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*User, error) {
	userDto, err := s.UserRepository.GetByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errhandler.NewNotFoundError("user not found")
		}

		return nil, errhandler.NewApplicationError(err.Error())
	}

	return &User{
		ID:    userDto.ID,
		Email: userDto.Email,
		Name:  userDto.Name,
	}, nil

}
