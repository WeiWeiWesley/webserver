package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//WsCommand websocket input
type WsCommand struct {
	Command string `json:"command"`
	Auth    string `json:"auth"`
	Data    string `json:"data"`
}

//ConnLock websocket lock
type ConnLock struct {
	Conn *websocket.Conn
	mu   *sync.Mutex
}

// Keep websocket connection
func wsKeep() gin.HandlerFunc {
	//移除origin 跨網域檢查
	var wsupgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	return func(c *gin.Context) {
		//升級連線，建立websocket連線
		wsConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("upgrader.Upgrade err: %s\n", err)
		}

		ws := ConnLock{
			Conn: wsConn,
			mu: &sync.Mutex{},
		}

		// 3min timeout
		idleLimit := 3 * time.Minute
		timeout := time.NewTimer(idleLimit)
		defer timeout.Stop()

		// Logout & close ws
		logout := make(chan bool, 1)
		defer close(logout)

		// Timeout
		go func() {
			for {
				select {
				case <-timeout.C:
					ws.writeMessage("timeout")
					logout <- true
					return
				}
			}
		}()

		//使用迴圈持續接收資料
		go func() {
			for {
				//讀取ws
				_, wsCommand, err := ws.Conn.ReadMessage()
				if err != nil {
					logout <- true
					return
				}

				// Reset timer
				timeout.Reset(idleLimit)

				var command WsCommand
				err = json.Unmarshal(wsCommand, &command)
				if err != nil {
					fmt.Printf("%+v \n", err)
					continue
				}

				err = ws.commandSwitch(&command)
				if err != nil {
					fmt.Printf("%+v \n", err)
				}

			}
		}()

		// Wating Exit
		select {
		case <-logout:
			time.Sleep(500)
			ws.Conn.Close()
		}

		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("cliententer panic: %v", p)
			}
			ws.Conn.Close()
		}()
	}
}

// Send message to client
func (ws *ConnLock) writeMessage(data string) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	err := ws.Conn.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		fmt.Printf("WriteMessage err: %s\nData: %s", err, data) //TODO 這個log應該要統一處理
	}

	return err
}
