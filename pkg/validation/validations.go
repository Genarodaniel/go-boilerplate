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

func IsUUID(id string) bool {
	if err := uuid.Validate(id); err != nil {
		return false
	}
	return true
}
