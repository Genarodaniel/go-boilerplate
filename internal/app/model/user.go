package model

import "time"

type PostUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PostUserRequest struct {
	Email string `validate:"required" json:"email"`
	Name  string `validate:"required" json:"name"`
}

type PostUserLogin struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserQueue struct {
	ClientID string `json:"client_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type AuthResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}
