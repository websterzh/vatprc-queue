package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vatprc-queue/gin/errors"
	"vatprc-queue/gin/services"
)

func CreateToken(c *gin.Context) (interface{}, error) {
	newToken := services.CreateToken("*")
	return newToken, nil
}

type DeleteTokenRequest struct {
	Token string `json:"token"`
}

func DeleteToken(c *gin.Context) (interface{}, error) {
	token, ok := c.Params.Get("token")
	if token == "" || !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: false,
			Message:          "Token is required",
		}
	}
	services.DeleteToken(token)
	return nil, nil
}
