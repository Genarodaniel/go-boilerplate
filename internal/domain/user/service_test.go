package user

import (
	"errors"
	repositoryMock "go-boilerplate/internal/repository/mock"
	"go-boilerplate/services/kafka"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	requestSuccess := &PostUserRequest{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
	}
	t.Run("should return an uuid when created a new user", func(t *testing.T) {
		userService := NewUserService(kafka.KafkaMock{}, repositoryMock.UserRepositoryMock{
			SaveUserResponse: gofakeit.UUID(),
		})
		response, err := userService.PostUser(ctx, requestSuccess)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Nil(t, uuid.Validate(response.UserID))
	})

	t.Run("should return an error when calling db to create order", func(t *testing.T) {
		kafkaMock := kafka.KafkaMock{}
		userRepositoryMock := repositoryMock.UserRepositoryMock{
			SaveUserError: errors.New("error name is too long"),
		}

		userService := NewUserService(kafkaMock, userRepositoryMock)
		response, err := userService.PostUser(ctx, requestSuccess)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, err, userRepositoryMock.SaveUserError)
	})

	t.Run("should return an error when calling kafka producer to create user", func(t *testing.T) {
		kafkaMock := kafka.KafkaMock{
			ProduceError: errors.New("error to conect to kafka"),
		}

		userRepositoryMock := repositoryMock.UserRepositoryMock{
			SaveUserResponse: gofakeit.UUID(),
		}

		userService := NewUserService(kafkaMock, userRepositoryMock)
		response, err := userService.PostUser(ctx, requestSuccess)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error to conect to kafka")
	})

}
