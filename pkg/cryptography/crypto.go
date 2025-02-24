package cryptography

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	ClientSecret string
	ClientID     string
}

func GenerateOAuthSecrets() (Client, error) {
	client := Client{}
	var err error
	client.ClientID = generateClientID()
	client.ClientSecret, err = generateClientSecret()
	if err != nil {
		return client, err
	}

	return client, nil
}

func generateClientID() string {
	return uuid.NewString()
}

func generateClientSecret() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func HashSecret(secret string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	return string(hashed), err
}

func ValidateClientSecret(providedSecret, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedSecret))
	return err == nil
}
