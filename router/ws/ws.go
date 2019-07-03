package ws

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
func (ws *ConnLock) commandSwitch(command *ReqCmd) error {
	var err error

	switch command.Command {
	case "ping": //Ping
		pong := make(map[string]interface{})
		pong["command"] = command.Command
		pong["resp_time"] = time.Now()

		res, err := json.Marshal(pong)
		if err != nil {
			log.Println("JSON error:", err.Error())
			break
		}
		ws.writeMessage(string(res))

	default:
		ws.writeMessage("Command not exists")
		return errors.New("Command \"" + command.Command + "\" not exists")
	}

	return err
}
