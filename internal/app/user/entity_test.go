package user

import (
	"go-boilerplate/internal/app/model"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when name is empty", func(t *testing.T) {
		request := User{}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, model.ErrNameRequired, err)
	})

	t.Run("Should return an error when email is empty", func(t *testing.T) {
		request := User{
			Name: gofakeit.Name(),
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, model.ErrEmailRequired, err)
	})

	t.Run("Should return an error when email is invalid", func(t *testing.T) {
		request := User{
			Name:  gofakeit.Name(),
			Email: "not valid email",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, model.ErrEmailInvalid, err)
	})

	t.Run("Should return success when have an id but it's not an uuid", func(t *testing.T) {
		request := User{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			ID:    "not valid uuid",
		}
		err := request.Validate()

		assert.Equal(t, model.ErrInvalidUUID, err)
	})

	t.Run("Should return success when the request is valid", func(t *testing.T) {
		request := User{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			ID:    gofakeit.UUID(),
		}
		err := request.Validate()

		assert.Nil(t, err)
	})

}
