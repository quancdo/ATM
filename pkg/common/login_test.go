package common_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ATM/pkg/common"
	"github.com/ATM/pkg/reststructs"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	var login reststructs.Login
	reqBody := io.NopCloser(strings.NewReader("{\"api_key\":\"123\",\"pin\":\"345\"}"))
	err := common.ParseRequest(&http.Request{Body: reqBody}, &login)
	assert.Nil(t, err)
	assert.Equal(t, reststructs.Login{APIKey: "123", Pin: "345"}, login)
}
