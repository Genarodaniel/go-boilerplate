package user

import (
	"go-boilerplate/internal/app/model"
	"go-boilerplate/pkg/validation"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (user User) Validate() error {
	if err := user.ValidateName(); err != nil {
		return err
	}

	if err := user.ValidateEmail(); err != nil {
		return err
	}

	if err := user.ValidatePassword(); err != nil {
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

func (user *User) ValidatePassword() error {
	if len(user.Password) == 0 {
		return model.ErrRequiredPassword
	}

	if len(user.Password) > 64 || len(user.Password) < 8 {
		return model.ErrInvalidPasswordLength
	}

	if valid := validation.ValidatePassword(user.Password); !valid {
		return model.ErrInvalidPassword
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
