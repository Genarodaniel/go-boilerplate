package user

import (
	"go-boilerplate/internal/app/model"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		Expected    error
		TestMessage string
		UserParam   User
	}{
		{
			TestMessage: "Should return an error to validate name",
			Expected:    model.ErrRequiredName,
			UserParam:   User{},
		},
		{
			TestMessage: "Should return an error to validate email",
			Expected:    model.ErrRequiredEmail,
			UserParam: User{
				Name: gofakeit.Name(),
			},
		},
		{
			TestMessage: "Should return an error when have an id but it's not a UUID",
			Expected:    model.ErrInvalidUUID,
			UserParam: User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				ID:    "not valid uuid",
			},
		},
		{
			TestMessage: "Should return success when the entity is valid",
			Expected:    nil,
			UserParam: User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				ID:    gofakeit.UUID(),
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			err := testCase.UserParam.Validate()
			assert.Equal(t, testCase.Expected, err)
		})
	}

}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		Expected    error
		TestMessage string
		UserParam   User
	}{
		{
			TestMessage: "Should return an error when email is empty",
			Expected:    model.ErrRequiredEmail,
			UserParam: User{
				Name: gofakeit.Name(),
			},
		},

		{
			TestMessage: "Should return an error when email is invalid",
			Expected:    model.ErrInvalidEmail,
			UserParam: User{
				Name:  gofakeit.Name(),
				Email: "not valid email",
			},
		},
		{
			TestMessage: "Should return nil when email is valid",
			Expected:    nil,
			UserParam: User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			err := testCase.UserParam.ValidateEmail()
			assert.Equal(t, testCase.Expected, err)
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		Expected    error
		TestMessage string
		UserParam   User
	}{
		{
			TestMessage: "Should return an error when name is empty",
			Expected:    model.ErrRequiredName,
			UserParam:   User{},
		},
		{
			TestMessage: "Should return an error when name is more than 120 chars",
			Expected:    model.ErrInvalidName,
			UserParam: User{
				Name: "this is a full name of a person that has more than one hundred twenty characters its hard to reach that number but its for test",
			},
		},

		{
			TestMessage: "Should not return error when name is valid",
			Expected:    nil,
			UserParam: User{
				Name: "Daniel silva genaro",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			err := testCase.UserParam.ValidateName()
			assert.Equal(t, testCase.Expected, err)
		})
	}
}
