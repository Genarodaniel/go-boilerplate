package validation

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var ErrInvalidUUID = errors.New("invalid uuid")

func ValidateEmail(email string) bool {
	regex := `^[a-z0-9!#$%&'*+/=?^_` + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + `{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9]){1,4}?$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func ValidateUUID(id string) bool {
	if err := uuid.Validate(id); err != nil {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`\d`)
	special := regexp.MustCompile(`[\W_]`) // Matches special characters (includes `_`)

	return lowercase.MatchString(password) &&
		uppercase.MatchString(password) &&
		number.MatchString(password) &&
		special.MatchString(password)
}
