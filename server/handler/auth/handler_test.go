package auth

import (
	"bytes"
	"encoding/json"
	"go-boilerplate/internal/app/auth"
	"go-boilerplate/internal/app/model"
	"go-boilerplate/internal/infra/mock"
	"go-boilerplate/pkg/customerror"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	repositoryMock "go-boilerplate/internal/repository/mock"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repositoryMock := repositoryMock.UserRepositoryMock{}
	authService := auth.NewAuthService(repositoryMock)

	router := gin.Default()
	authHandler := NewAuthHandler(authService)
	router.POST("/v1/auth", authHandler.HandleAuth)

	path := "/v1/auth"

	t.Run("Should return error for invalid JSON payload", func(t *testing.T) {
		mockRequest := map[string]interface{}{
			"client_id":     123,            // Invalid type (should be string)
			"client_secret": []int{1, 2, 3}, // Invalid type (should be string)
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioRequest := io.NopCloser(bytes.NewBuffer(requestBytes))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		authHandler.HandleAuth(ctx)
		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), model.ErrInvalidRequest.Error())
	})

	t.Run("Should return validation error for missing credentials", func(t *testing.T) {
		mockRequest := model.OAuthRequest{
			ClientID:     "",
			ClientSecret: "",
			GrantType:    "client_credentials",
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioRequest := io.NopCloser(bytes.NewBuffer(requestBytes))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		authHandler.HandleAuth(ctx)
		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, string(response), model.ErrRequiredCredentials.Error())
	})

	t.Run("Should return error for unsupported grant type", func(t *testing.T) {
		mockRequest := model.OAuthRequest{
			ClientID:     gofakeit.UUID(),
			ClientSecret: "super_secret",
			GrantType:    "password", // Unsupported grant type for this flow
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioRequest := io.NopCloser(bytes.NewBuffer(requestBytes))

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		authHandler.HandleAuth(ctx)
		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), model.ErrInvalidGrantType.Error())
	})

	t.Run("Should return unauthorized error for invalid credentials", func(t *testing.T) {
		mockRequest := model.OAuthRequest{
			ClientID:     gofakeit.UUID(),
			ClientSecret: "invalid_secret",
			GrantType:    "client_credentials",
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioRequest := io.NopCloser(bytes.NewBuffer(requestBytes))

		authServiceMock := mock.AuthServiceMock{
			AuthenticateResponse: auth.Auth{},
			AuthenticateError:    customerror.NewUnauthorizedError(model.ErrInvalidCredentials.Error()),
		}
		authHandler := NewAuthHandler(authServiceMock)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		authHandler.HandleAuth(ctx)
		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, string(response), authServiceMock.AuthenticateError.Error())
	})

	t.Run("Should return the OAuth token response", func(t *testing.T) {
		mockRequest := model.OAuthRequest{
			ClientID:     gofakeit.UUID(),
			ClientSecret: "valid_secret",
			GrantType:    "client_credentials",
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioRequest := io.NopCloser(bytes.NewBuffer(requestBytes))

		authServiceMock := mock.AuthServiceMock{
			AuthenticateResponse: auth.Auth{
				AccessToken:  "generated_access_token",
				RefreshToken: "generated_refresh_token",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				ExpiresAt:    time.Now().Add(3600 * time.Second),
			},
			AuthenticateError: nil,
		}
		authHandler := NewAuthHandler(authServiceMock)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		authHandler.HandleAuth(ctx)
		response, _ := io.ReadAll(w.Body)

		responseServiceBytes, _ := json.Marshal(authServiceMock.AuthenticateResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(responseServiceBytes), string(response))
	})
}
