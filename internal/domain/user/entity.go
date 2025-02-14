package user

import (
	"errors"
	"go-boilerplate/pkg/validation"

	"github.com/google/uuid"
)

type UserStatus string

const (
	UserStatusCreated UserStatus = "created"
)

type PostUserResponse struct {
	UserID string `json:"user_id"`
}

type PostUserRequest struct {
	ClientID          string  `json:"client_id"`
	StoreID           string  `json:"store_id"`
	NotificationEmail string  `json:"notification_email"`
	Status            string  `json:"status"`
	UserID            string  `json:"user_id"`
	Amount            float64 `json:"amount"`
}

func (request *PostUserRequest) Validate() error {
	if request.Amount <= 0 {
		return errors.New("amount must be a positive number")
	}

	if err := uuid.Validate(request.ClientID); err != nil {
		return errors.New("client_id must be a uuid")
	}

	if err := uuid.Validate(request.StoreID); err != nil {
		return errors.New("store_id must be a uuid")
	}

	if valid := validation.ValidateEmail(request.NotificationEmail); !valid {
		return errors.New("notification_email invalid")
	}

	return nil
}
