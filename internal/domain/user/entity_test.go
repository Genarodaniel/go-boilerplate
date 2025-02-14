package user_test

import (
	"go-boilerplate/internal/domain/user"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when amount is empty", func(t *testing.T) {
		request := user.PostUserRequest{}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "amount must be a positive number", err.Error())
	})

	t.Run("Should return an error when amount is negative", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount: -123.00,
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "amount must be a positive number", err.Error())
	})

	t.Run("Should return an error when clientID is empty", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount: 123.00,
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "client_id must be a uuid", err.Error())
	})

	t.Run("Should return an error when clientID is not a uuid", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount:   123.00,
			ClientID: "not a uuid",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "client_id must be a uuid", err.Error())
	})

	t.Run("Should return an error when storeID is empty", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "store_id must be a uuid", err.Error())
	})

	t.Run("Should return an error when storeID is not a uuid", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  "not a uuid",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "store_id must be a uuid", err.Error())
	})

	t.Run("Should return an error when email is invalid", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: "not a email",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "notification_email invalid", err.Error())
	})

	t.Run("Should return success when the request is valid", func(t *testing.T) {
		request := user.PostUserRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: gofakeit.Email(),
		}
		err := request.Validate()

		assert.Nil(t, err)
	})

}
