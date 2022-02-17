package handlers_test

import (
	"testing"

	"github.com/ATM/pkg/handlers"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {

	h := &handlers.ATMHandler{}
	c, resp := createTestContext("POST", map[string]string{"api_key": "john_apikey", "pin": "1234"}, "")
	h.LoginWithPin(c)
	assert.Equal(t, 200, resp.Code)
}
