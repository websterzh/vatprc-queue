package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vatprc-queue/gin/errors"
	"vatprc-queue/gin/services"
)

type UpdateStatusRequest struct {
	Cid    string `json:"cid"`
	Status int    `json:"status"`
}

func UpdateStatus(c *gin.Context) (interface{}, error) {
	airport, ok := c.Params.Get("airport")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "airport parameter is required",
		}
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: false,
			Message:          err.Error(),
		}
	}

	err := services.UpdateStatus(airport, req.Cid, req.Status)
	if err != nil {
		return nil, errors.ApiError{
			Status:           http.StatusInternalServerError,
			Code:             http.StatusInternalServerError,
			ShowInProduction: false,
			Message:          err.Error(),
		}
	}
	return services.GetQueueResult(airport, true), nil
}

type UpdateOrderRequest struct {
	Cid    string `json:"cid"`
	Before string `json:"before"`
}

func UpdateOrderHandler(c *gin.Context) (interface{}, error) {
	airport, ok := c.Params.Get("airport")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "airport parameter is required",
		}
	}
	var req UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: false,
			Message:          err.Error(),
		}
	}

	err := services.UpdateOrder(airport, req.Cid, req.Before)
	if err != nil {
		return nil, errors.ApiError{
			Status:           http.StatusInternalServerError,
			Code:             http.StatusInternalServerError,
			ShowInProduction: false,
			Message:          err.Error(),
		}
	}
	return services.GetQueueResult(airport, true), nil
}

func GetQueueHandler(c *gin.Context) (interface{}, error) {
	airport, ok := c.Params.Get("airport")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "airport parameter is required",
		}
	}
	return services.GetQueueResult(airport, true), nil
}
