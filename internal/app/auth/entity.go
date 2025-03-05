package auth

import "time"

type Auth struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int64
	ExpiresAt    time.Time
}
