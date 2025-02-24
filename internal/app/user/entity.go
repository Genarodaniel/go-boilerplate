package user

import (
	"go-boilerplate/internal/app/model"
	"go-boilerplate/pkg/cryptography"
	"go-boilerplate/pkg/validation"

	"github.com/google/uuid"
)

type UserInterface interface {
	Validate() error
	ValidateName() error
	ValidateEmail() error
}
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewUser(email, name string) (*User, error) {
	client, err := cryptography.GenerateOAuthSecrets()
	if err != nil {
		return nil, err
	}

	user := User{
		ID:           uuid.NewString(),
		Name:         name,
		Email:        email,
		ClientID:     client.ClientID,
		ClientSecret: client.ClientSecret,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil

}

func (user *User) Validate() error {
	if err := user.ValidateName(); err != nil {
		return err
	}

	if err := user.ValidateEmail(); err != nil {
		return err
	}

	if user.ID != "" {
		if valid := validation.ValidateUUID(user.ID); !valid {
			return model.ErrInvalidUUID
		}
	}

	return nil
}

func (user *User) ValidateName() error {
	if len(user.Name) == 0 {
		return model.ErrRequiredName
	}

	if len(user.Name) > 120 {
		return model.ErrInvalidName
	}

	return nil
}

func (user *User) ValidateEmail() error {
	if len(user.Email) == 0 {
		return model.ErrRequiredEmail
	}

	if valid := validation.ValidateEmail(user.Email); !valid {
		return model.ErrInvalidEmail
	}

	return nil
}
