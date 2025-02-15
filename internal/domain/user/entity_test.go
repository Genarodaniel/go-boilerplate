package user_test

import (
	"go-boilerplate/internal/domain/user"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when name is empty", func(t *testing.T) {
		request := user.PostUserRequest{}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, user.ErrNameRequired, err)
	})

	t.Run("Should return an error when email is empty", func(t *testing.T) {
		request := user.PostUserRequest{
			Name: gofakeit.Name(),
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, user.ErrEmailRequired, err)
	})

	t.Run("Should return an error when email is invalid", func(t *testing.T) {
		request := user.PostUserRequest{
			Name:  gofakeit.Name(),
			Email: "not valid email",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, user.ErrEmailInvalid, err)
	})

	t.Run("Should return success when the request is valid", func(t *testing.T) {
		request := user.PostUserRequest{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		}
		err := request.Validate()

		assert.Nil(t, err)
	})

}
