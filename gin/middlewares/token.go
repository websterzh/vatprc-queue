package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
	}
}
