package middleware

import "github.com/gin-gonic/gin"

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
