package handlers_test

import (
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func createTestContext(method string, params map[string]string, data string) (c *gin.Context, resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()

	c, _ = gin.CreateTestContext(resp)
	c.Request = httptest.NewRequest(method, "/", strings.NewReader(data))
	for k, v := range params {
		c.Params = append(c.Params, gin.Param{Key: k, Value: v})
	}
	return c, resp
}
