package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		api_key := c.Param("api_key")
		pin := c.Param("pin")

		if len(api_key) == 0 {
			log.Print("missing apikey")
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{"error_msg": "internal_error"})
			return
		}

		if len(pin) == 0 {
			log.Print("missing pin")
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{"error_msg": "pin is missing in the url parameter"})
			return
		}
		c.Next()
	}
}
