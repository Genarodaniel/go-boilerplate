package user

import (
	"database/sql"
	"errors"
	"go-boilerplate/internal/app/model"
	repositoryMock "go-boilerplate/internal/repository/mock"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/internal/services/kafka"
	"net/http/httptest"
	"testing"

	"go-boilerplate/pkg/customerror"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	requestSuccess := User{
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: "Aa1!abcd",
	}
	t.Run("should return an validation error", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{
			SaveUserResponse: gofakeit.UUID(),
		})
		response, err := userService.Register(ctx, requestSuccess)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Nil(t, uuid.Validate(response.ID))
	})

	t.Run("should return an uuid when created a new user", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{})
		requestWithEmailError := User{
			Name:  gofakeit.Name(),
			Email: "not valid email",
		}
		response, err := userService.Register(ctx, requestWithEmailError)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, customerror.NewValidationError(model.ErrInvalidEmail.Error()), err)
	})

	t.Run("should return an error when calling db to create order", func(t *testing.T) {
		kafkaMock := kafka.KafkaMock{}
		userRepositoryMock := repositoryMock.UserRepositoryMock{
			SaveUserError: errors.New("error name is too long"),
		}

		userService := NewUserService(kafkaMock, userRepositoryMock)
		response, err := userService.Register(ctx, requestSuccess)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, customerror.NewApplicationError(userRepositoryMock.SaveUserError.Error()), err)
	})

	t.Run("should return an error when calling kafka producer to create user", func(t *testing.T) {
		kafkaMock := kafka.KafkaMock{
			ProduceError: errors.New("error to conect to kafka"),
		}

		userRepositoryMock := repositoryMock.UserRepositoryMock{
			SaveUserResponse: gofakeit.UUID(),
		}

		userService := NewUserService(kafkaMock, userRepositoryMock)
		response, err := userService.Register(ctx, requestSuccess)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, customerror.NewApplicationError(kafkaMock.ProduceError.Error()), err)
	})

}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	userID := gofakeit.UUID()
	userDto := userRepository.UserDTO{
		ID:    userID,
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
	}

	t.Run("should return a user when user exists", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{
			GetUserResponse: userDto,
		})
		response, err := userService.Get(ctx, userID)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, userID, response.ID)
	})

	t.Run("should return an error when user does not exist", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{
			GetUserError: sql.ErrNoRows,
		})
		response, err := userService.Get(ctx, userID)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualError(t, customerror.NewNotFoundError("user not found"), err.Error())
	})

	t.Run("should return an error when user does not exist", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{
			GetUserError: errors.New("sql error"),
		})
		response, err := userService.Get(ctx, userID)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualError(t, customerror.NewApplicationError("sql error"), err.Error())
	})
}
