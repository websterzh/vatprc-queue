package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  errorToString(r),
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
