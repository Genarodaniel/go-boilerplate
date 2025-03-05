package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/app/user"
	"go-boilerplate/internal/infra/mock"
	repositoryMock "go-boilerplate/internal/repository/mock"
	"go-boilerplate/internal/services/kafka"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	kafkaMock := kafka.KafkaMock{}
	repositoryMock := repositoryMock.UserRepositoryMock{}
	userService := user.NewUserService(kafkaMock, repositoryMock)
	Router(&gin.Default().RouterGroup, userService)
	path := "/v1/user"

	t.Run("Should return error when payload is empty", func(t *testing.T) {
		userService := user.NewUserService(kafkaMock, repositoryMock)
		userHandler := NewUserHandler(userService)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, nil)
		userHandler.HandleRegister(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return error when the given params are of different types than the expected", func(t *testing.T) {
		mockRequest := map[string]interface{}{
			"name": 123,
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userService := user.NewUserService(kafkaMock, repositoryMock)
		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandleRegister(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "cannot unmarshal")
	})

	t.Run("Should return a validation error", func(t *testing.T) {
		mockRequest := model.PostUserRequest{
			Name:  "",
			Email: "",
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userService := user.NewUserService(kafkaMock, repositoryMock)
		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandleRegister(ctx)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("Should return an service error", func(t *testing.T) {
		errorMessage := "error to save user"
		mockRequest := model.PostUserRequest{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		}

		userService := mock.UserServiceMock{
			UserResponse:  user.User{},
			RegisterError: errors.New(errorMessage),
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		userHandler := NewUserHandler(userService)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandleRegister(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, string(response), errorMessage)
	})

	t.Run("Should create the user", func(t *testing.T) {
		mockRequest := model.PostUserRequest{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		}

		userService := mock.UserServiceMock{
			UserResponse: user.User{
				ID: uuid.NewString(),
			},
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandleRegister(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, string(response), "id")
	})
}

func TestHandleGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	kafkaMock := kafka.KafkaMock{}
	repositoryMock := repositoryMock.UserRepositoryMock{}
	userService := user.NewUserService(kafkaMock, repositoryMock)
	Router(&gin.Default().RouterGroup, userService)
	path := "/user/v1/:id"

	t.Run("Should return error when user ID is invalid", func(t *testing.T) {
		userService := user.NewUserService(kafkaMock, repositoryMock)
		userHandler := NewUserHandler(userService)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = gin.Params{gin.Param{Key: "id", Value: "invalid-uuid"}}
		userHandler.HandleGet(ctx)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("Should return error when user not found", func(t *testing.T) {
		userService := mock.UserServiceMock{
			GetError: errors.New("user not found"),
		}
		userHandler := NewUserHandler(userService)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = gin.Params{gin.Param{Key: "id", Value: uuid.NewString()}}
		userHandler.HandleGet(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, string(response), "user not found")
	})

	t.Run("Should return user details", func(t *testing.T) {
		userID := uuid.NewString()
		userService := mock.UserServiceMock{
			UserResponse: user.User{
				ID:    userID,
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			},
		}
		userHandler := NewUserHandler(userService)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = gin.Params{gin.Param{Key: "id", Value: userID}}
		userHandler.HandleGet(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, string(response), userID)
	})
}
