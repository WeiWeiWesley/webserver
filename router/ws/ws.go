package ws

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//StartWsRouter StartWsRouter
func StartWsRouter() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	ws := r.Group("/ws") //Websocket API

	// Websocket API
	ws.GET("/keep", wsKeep())

	r.Run(":3900")
}

// WS API command router
func commandSwitch(ws *websocket.Conn, command *WsCommand) error {
	var (
		err error
	)

	switch command.Command {
	case "ping": //Ping
		writeMessage(ws, "pong")

	default:
		writeMessage(ws, "Command not exists")
		return errors.New("Command \"" + command.Command + "\" not exists")
	}

	return err
}
