package user

import (
	"context"
	"go-boilerplate/internal/repository"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/services/kafka"
)

type UserServiceInterface interface {
	PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error)
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

func (s *UserService) PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error) {
	userDto := userRepository.User{
		Name:  user.Name,
		Email: user.Email,
	}

	userID, err := s.UserRepository.Save(ctx, userDto)
	if err != nil {
		return nil, err
	}

	if err := s.KafkaProducer.Produce(ctx, "users", "user.create", userDto); err != nil {
		return nil, err
	}

	return &PostUserResponse{
		UserID: userID,
	}, nil
}
