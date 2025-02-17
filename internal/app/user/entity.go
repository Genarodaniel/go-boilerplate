package user

import (
	"errors"
	"go-boilerplate/pkg/validation"
)

type PostUserResponse struct {
	UserID string `json:"user_id"`
}

type PostUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

var ErrNameRequired error = errors.New("name is required")
var ErrEmailRequired error = errors.New("email is required")
var ErrEmailInvalid error = errors.New("email is invalid")

func (user *PostUserRequest) Validate() error {
	if len(user.Name) == 0 {
		return ErrNameRequired
	}

	if len(user.Email) == 0 {
		return ErrEmailRequired
	}

	if valid := validation.ValidateEmail(user.Email); !valid {
		return ErrEmailInvalid
	}

	return nil
}
