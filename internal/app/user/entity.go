package user

import (
	"go-boilerplate/internal/app/model"
	"go-boilerplate/pkg/validation"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (user User) Validate() error {
	if len(user.Name) == 0 {
		return model.ErrNameRequired
	}

	if len(user.Email) == 0 {
		return model.ErrEmailRequired
	}

	if valid := validation.ValidateEmail(user.Email); !valid {
		return model.ErrEmailInvalid
	}

	if user.ID != "" {
		if valid := validation.IsUUID(user.ID); !valid {
			return model.ErrInvalidUUID
		}
	}

	return nil
}
