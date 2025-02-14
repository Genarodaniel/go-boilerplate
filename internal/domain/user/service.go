package user

import (
	"context"
	"go-boilerplate/services/kafka"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error)
}

type UserService struct {
	KafkaProducer kafka.KafkaInterface
}

func NewuserService(kafkaProducer kafka.KafkaInterface) *UserService {
	return &UserService{
		KafkaProducer: kafkaProducer,
	}
}

func (s *UserService) PostUser(ctx context.Context, user *PostUserRequest) (*PostUserResponse, error) {
	user.UserID = uuid.NewString()
	user.Status = string(UserStatusCreated)
	if err := s.KafkaProducer.Produce(ctx, "users", "user.create", user); err != nil {
		return nil, err
	}

	return &PostUserResponse{
		UserID: user.UserID,
	}, nil
}
