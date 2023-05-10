package handlers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(_ *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"message": "HELLO WORLD!",
	}, nil
}
