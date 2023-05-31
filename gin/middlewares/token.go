package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vatprc-queue/config"
	"vatprc-queue/gin/services"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		if !services.HasToken(token, "*") {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AtcCenterAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		if token != config.File.Section("app").Key("token_generate_key").String() {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
