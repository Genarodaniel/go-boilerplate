package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	type validationCase struct {
		Email       string
		Valid       bool
		TestMessage string
	}
	cases := []validationCase{
		{Email: "danielgenaro@gmail.com", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "danielgenaro@hotmail.com", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "danielgenaro", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "123132WW0090[[[]]]ss", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "teste@teste@teste.com", Valid: false, TestMessage: "Should return false with an invalid email"},
		{Email: "daniel.genaro@domain.com.br", Valid: true, TestMessage: "Should return true with the valid email"},
		{Email: "daniel.genaro@domain.co", Valid: true, TestMessage: "Should return true with the valid email"},
	}

	for _, testcase := range cases {
		t.Run(testcase.TestMessage, func(t *testing.T) {
			valid := ValidateEmail(testcase.Email)
			assert.Equal(t, testcase.Valid, valid)
		})
	}

}
