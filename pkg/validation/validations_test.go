package validation

import (
	"testing"

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
