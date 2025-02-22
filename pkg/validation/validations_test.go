package validation

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		Email       string
		Valid       bool
		TestMessage string
	}{
		{Email: "danielgenaro@gmail.com", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "danielgenaro@hotmail.com", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "danielgenaro", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "123132WW0090[[[]]]ss", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "teste@teste@teste.com", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "daniel.genaro@domain.com.br", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "daniel.genaro@domain.co", Valid: true, TestMessage: "Should return true with the valid email"},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			valid := ValidateEmail(testCase.Email)
			assert.Equal(t, testCase.Valid, valid)
		})
	}
}

func TestValidateValidateUUID(t *testing.T) {
	tests := []struct {
		UUID        string
		Valid       bool
		TestMessage string
	}{
		{UUID: uuid.NewString(), Valid: true, TestMessage: "Should return true with the valid uuid"},
		{UUID: "29c89359-d301-4dd8-9e4c-e1d00e6d7ae3", Valid: true, TestMessage: "Should return true with the valid uuid v4"},
		{UUID: "danielgenaro", Valid: false, TestMessage: "Should return false with an invalid uuid"},
		{UUID: "123132WW0090[[[]]]ss", Valid: false, TestMessage: "Should return false with an invalid uuid"},
		{UUID: "teste@teste@teste.com", Valid: false, TestMessage: "Should return false with an invalid uuid"},
		{UUID: "01951b47-c00d-7bdb-9563-eee948ee39ed", Valid: true, TestMessage: "Should return true with the valid uuid v7"},
		{UUID: "c5cc5b32-ee4b-11ef-9cd2-0242ac120002", Valid: true, TestMessage: "Should return true with the valid uuid v1"},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			valid := ValidateUUID(testCase.UUID)
			assert.Equal(t, testCase.Valid, valid)
		})
	}
}

func TestValidatePassword(t *testing.T) {
	type validationCase struct {
		Password    string
		Valid       bool
		TestMessage string
	}
	tests := []validationCase{
		// Combinations missing one required type

		{Password: "", Valid: false, TestMessage: "Should return false with empty password"},
		{Password: gofakeit.Password(true, false, false, false, false, 5), Valid: false, TestMessage: "Should return false when password contains only lower case letters"},
		{Password: gofakeit.Password(false, true, false, false, false, 5), Valid: false, TestMessage: "Should return false when password contains only upper case letters"},
		{Password: gofakeit.Password(false, false, true, false, false, 5), Valid: false, TestMessage: "Should return false when password contains only numbers"},
		{Password: gofakeit.Password(false, false, false, true, false, 5), Valid: false, TestMessage: "Should return false when password contains only special characters"},

		// Edge Cases: Missing one or more required character types
		{Password: gofakeit.Password(true, false, true, false, false, 8), Valid: false, TestMessage: "Should return false when password contains only numbers and lower case letters"},
		{Password: gofakeit.Password(true, true, false, false, false, 8), Valid: false, TestMessage: "Should return false when password contains only upper and lower case letters"},
		{Password: gofakeit.Password(false, true, true, false, false, 8), Valid: false, TestMessage: "Should return false when password contains only numbers and upper case letters"},
		{Password: gofakeit.Password(false, false, true, true, false, 8), Valid: false, TestMessage: "Should return false when password contains only numbers and special characters"},
		{Password: gofakeit.Password(true, false, false, true, false, 8), Valid: false, TestMessage: "Should return false when password contains only lower case letters and special characters"},
		{Password: gofakeit.Password(false, true, false, true, false, 8), Valid: false, TestMessage: "Should return false when password contains only upper case letters and special characters"},

		// Valid Cases: Passwords with all required components
		{Password: "Aa1!abcd", Valid: true, TestMessage: "Should return true when password contains uppercase, lowercase, number, and special character"},
		{Password: "XyZ@2024", Valid: true, TestMessage: "Should return true when password contains all required characters"},
	}

	for _, testCase := range tests {
		t.Run(testCase.TestMessage, func(t *testing.T) {
			valid := ValidatePassword(testCase.Password)
			assert.Equal(t, testCase.Valid, valid)
		})
	}

}
