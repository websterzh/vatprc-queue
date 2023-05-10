package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"nhooyr.io/websocket"
	"vatprc-queue/sockets"
)

func NewWSConnection(c *gin.Context) {
	airport, ok := c.Params.Get("airport")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("airport parameter is required"))
	}
	// FIXME: CORS
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	sockets.AddClient(conn, airport)
}
