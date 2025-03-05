package model

import "time"

type OAuthResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type OAuthRequest struct {
	ClientID     string `validate:"required" json:"client_id"`
	ClientSecret string `validate:"required" json:"client_secret"`
	GrantType    string `validate:"required" json:"grant_type"`
}
