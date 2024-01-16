package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"vatprc-queue/gin/errors"
	"vatprc-queue/gin/handlers"
	"vatprc-queue/gin/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	v1 := router.Group("v1")
	{
		v1.GET("/hello", errors.ErrorWrapper(handlers.HelloWorld))
		v1.GET("/queue", errors.ErrorWrapper(handlers.GetMultipleQueuesHandler))
		token := v1.Group("/token")
		{
			token.Use(middlewares.AtcCenterAuth())
			token.POST("", errors.ErrorWrapper(handlers.CreateToken))
			token.DELETE("/:token", errors.ErrorWrapper(handlers.DeleteToken))
		}
		airport := v1.Group("/:airport")
		{
			protected := airport.Group("")
			protected.Use(middlewares.TokenAuth())
			{
				protected.PATCH("/status", errors.ErrorWrapper(handlers.UpdateStatus))
				protected.PATCH("/order", errors.ErrorWrapper(handlers.UpdateOrderHandler))
				protected.DELETE("/status/:callsign", errors.ErrorWrapper(handlers.DeleteStatus))
			}
			airport.GET("/queue", errors.ErrorWrapper(handlers.GetQueueHandler))
			airport.GET("/ws", handlers.NewWSConnection)
		}
	}

	return router
}
