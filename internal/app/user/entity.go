package user

import (
	"errors"
	userRepository "go-boilerplate/internal/repository/user"
	"go-boilerplate/pkg/validation"
)

var ErrNameRequired error = errors.New("name is required")
var ErrEmailRequired error = errors.New("email is required")
var ErrEmailInvalid error = errors.New("email is invalid")
var ErrInvalidUUID = errors.New("invalid user id")
var ErrRequiredUserID = errors.New("user id is required")

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PostUserResponse struct {
	UserID string `json:"user_id"`
}

type PostUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GetUserRequest string

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

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

func (user *GetUserRequest) Validate() error {
	if user == nil {
		return ErrRequiredUserID
	}

	if len(string(*user)) == 0 {
		return ErrRequiredUserID
	}

	if valid := validation.IsUUID(string(*user)); !valid {
		return ErrInvalidUUID
	}

	return nil
}

func (user *PostUserRequest) ToEntity() *User {
	return &User{
		Email: user.Email,
		Name:  user.Name,
	}
}

func (user User) ToEntity(dto userRepository.UserDTO) *User {
	return &User{
		ID:    dto.ID,
		Email: dto.Email,
		Name:  dto.Name,
	}
}
