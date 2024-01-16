package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vatprc-queue/gin/errors"
	"vatprc-queue/gin/services"
)

type UpdateStatusRequest struct {
	Callsign string `json:"callsign"`
	Status   int    `json:"status"`
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

	err := services.UpdateStatus(airport, req.Callsign, req.Status)
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
	Callsign string `json:"callsign"`
	Before   string `json:"before"`
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

	err := services.UpdateOrder(airport, req.Callsign, req.Before)
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

func GetMultipleQueuesHandler(c *gin.Context) (interface{}, error) {
	param, ok := c.Params.Get("airports")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "airports parameter is required",
		}
	}
	airports := strings.Split(param, ",")
	if len(airports) < 1 || len(airports) > 10 {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "enter 1 to 10 airports",
		}
	}
	result := make(map[string][]services.QueueResult)
	for _, airport := range airports {
		result[airport] = services.GetQueueResult(airport, true)
	}
	return result, nil
}

func DeleteStatus(c *gin.Context) (interface{}, error) {
	airport, ok := c.Params.Get("airport")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "airport parameter is required",
		}
	}

	callsign, ok := c.Params.Get("callsign")
	if !ok {
		return nil, errors.ApiError{
			Status:           http.StatusBadRequest,
			Code:             http.StatusBadRequest,
			ShowInProduction: true,
			Message:          "callsign parameter is required",
		}
	}

	services.RemoveFromQueue(airport, callsign)
	return services.GetQueueResult(airport, true), nil
}
