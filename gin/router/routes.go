package router

import (
	"github.com/gin-gonic/gin"
	"vatprc-queue/gin/errors"
	"vatprc-queue/gin/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("v1")
	{
		v1.GET("/hello", errors.ErrorWrapper(handlers.HelloWorld))
		airport := v1.Group("/:airport")
		{
			airport.PATCH("/status", errors.ErrorWrapper(handlers.UpdateStatus))
			airport.PATCH("/order", errors.ErrorWrapper(handlers.UpdateOrderHandler))
			airport.GET("/queue", errors.ErrorWrapper(handlers.GetQueueHandler))
			airport.GET("/ws", handlers.NewWSConnection)
		}
	}

	return router
}
