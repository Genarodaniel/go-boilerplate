package user

import (
	"errors"
	"go-boilerplate/services/kafka"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	t.Run("should return an uuid when created a new user", func(t *testing.T) {
		userService := NewuserService(kafka.KafkaSpy{})
		response, err := userService.PostUser(ctx, &PostUserRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		})

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Nil(t, uuid.Validate(response.UserID))
	})

	t.Run("should return an error when calling kafka producer to create user", func(t *testing.T) {
		userService := NewuserService(kafka.KafkaSpy{
			ProduceError: errors.New("error to conect to kafka"),
		})
		response, err := userService.PostUser(ctx, &PostUserRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		})

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error to conect to kafka")
	})

}
