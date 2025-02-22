package helper

import "github.com/golang-jwt/jwt"

type JWTHelper interface {
	GetClaims(token string) (*jwt.Claims, error)
	// GenerateJWT()
}

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) JWTHelper {
	return &JWT{
		SecretKey: secretKey,
	}
}

func (h *JWT) GetClaims(token string) (*jwt.Claims, error) {
	return nil, nil
}
