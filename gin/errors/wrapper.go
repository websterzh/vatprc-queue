package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vatprc-queue/config"
)

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			apiError := err.(ApiError)
			debug, err := config.File.Section("app").Key("debug").Bool()
			if err != nil {
				debug = true
			}
			if (!apiError.ShowInProduction) && (!debug) {
				apiError.Message = "An error occurs."
			}
			c.JSON(apiError.Status, apiError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
