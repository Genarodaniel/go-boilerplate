package user

import (
	"context"
	"database/sql"
	"go-boilerplate/internal/repository"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/internal/services/kafka"
	"go-boilerplate/pkg/customerror"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Register(ctx context.Context, user User) (*User, error)
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

func (s *UserService) Register(ctx context.Context, user User) (*User, error) {
	if err := user.Validate(); err != nil {
		return nil, customerror.NewValidationError(err.Error())
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	userDto := userRepository.UserDTO{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashPassword),
	}

	userID, err := s.UserRepository.Save(ctx, userDto)
	if err != nil {
		return nil, customerror.NewApplicationError(err.Error())
	}

	user.ID = userID
	if err := s.KafkaProducer.Produce(ctx, "users", "user.create", user); err != nil {
		return nil, customerror.NewApplicationError(err.Error())
	}

	return &user, nil
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

// func (s *UserService) Update(ctx context.Context, userID string) (*User, error) {
// 	userDto, err := s.UserRepository.GetByID(ctx, userID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, customerror.NewNotFoundError("user not found")
// 		}

// 		return nil, customerror.NewApplicationError(err.Error())
// 	}

// 	return &User{
// 		ID:    userDto.ID,
// 		Email: userDto.Email,
// 		Name:  userDto.Name,
// 	}, nil

// }
