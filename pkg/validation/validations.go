package validation

import "regexp"

func ValidateEmail(email string) bool {
	regex := `^[a-z0-9!#$%&'*+/=?^_` + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + `{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9]){1,4}?$`
	match, _ := regexp.MatchString(regex, email)
	return match
}
