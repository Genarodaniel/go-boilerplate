package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-boilerplate/services/kafka"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandlePostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	kafkaMock := kafka.KafkaSpy{}
	userService := NewuserService(kafkaMock)
	Router(&gin.Default().RouterGroup, userService)
	path := "/user/v1/"

	t.Run("Should return error when payload is empty", func(t *testing.T) {
		userService := NewuserService(kafkaMock)
		userHandler := NewUserHandler(userService)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, nil)
		userHandler.HandlePostUser(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return error when the given params are of different types than the expected", func(t *testing.T) {
		mockRequest := map[string]interface{}{
			"client_id": 123,
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userService := NewuserService(kafkaMock)
		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandlePostUser(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "cannot unmarshal")
	})

	t.Run("Should return a validation error", func(t *testing.T) {
		mockRequest := PostUserRequest{
			Amount:   -123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userService := NewuserService(kafkaMock)
		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandlePostUser(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "amount must be a positive number")
	})

	t.Run("Should return an service error", func(t *testing.T) {
		errorMessage := "error to save user transaction"
		mockRequest := PostUserRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: gofakeit.Email(),
		}

		userService := UserServiceSpy{
			PostUserResponse: PostUserResponse{},
			PostUserError:    errors.New(errorMessage),
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		userHandler := NewUserHandler(userService)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandlePostUser(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, string(response), errorMessage)
	})

	t.Run("Should create the user", func(t *testing.T) {
		mockRequest := PostUserRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: gofakeit.Email(),
		}

		userService := UserServiceSpy{
			PostUserResponse: PostUserResponse{
				uuid.NewString(),
			},
			PostUserError: nil,
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		userHandler := NewUserHandler(userService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		userHandler.HandlePostUser(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, string(response), "user_id")
	})
}
