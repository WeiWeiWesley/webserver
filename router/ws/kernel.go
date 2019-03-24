package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//GameServer 註冊連線資訊
type GameServer struct {
	PartnerID string          `json:"partner_id"`
	IP        string          `json:"ip"`
	Token     string          `json:"token"`
	Conn      *websocket.Conn `json:"conn"`
	TTL       *time.Timer     `json:"ttl"`
}

//WsCommand websocket input
type WsCommand struct {
	Command string `json:"command"`
	Auth    string `json:"auth"`
	Data    string `json:"data"`
}

//[DANGER] 已註冊連線清單
var (
	serverMap     map[string]GameServer
	addServerChan chan GameServer
	reConnServer  chan GameServer
	pushChan      chan Push
)

//Push 推送目標
type Push struct {
	Target *websocket.Conn
	Data   string
}


// Keep websocket connection
func wsKeep() gin.HandlerFunc {
	//移除origin 跨網域檢查
	var wsupgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	return func(c *gin.Context) {
		//升級連線，建立websocket連線
		ws, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Printf("upgrader.Upgrade err: %s\n", err)
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
					writeMessage(ws, "timeout")
					logout <- true
					return
				}
			}
		}()

		//使用迴圈持續接收資料
		go func() {
			for {
				//讀取ws
				_, wsCommand, err := ws.ReadMessage()
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

				err = commandSwitch(ws, &command)
				if err != nil {
					fmt.Printf("%+v \n", err)
				}

			}
		}()

		// Wating Exit
		select {
		case <-logout:
			time.Sleep(500)
			ws.Close()
		}

		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("cliententer panic: %v", p)
			}
			ws.Close()
		}()
	}
}

// Send message to client
func writeMessage(ws *websocket.Conn, data string) error {
	err := ws.WriteMessage(websocket.TextMessage, []byte(data))
	if err != nil {
		fmt.Printf("WriteMessage err: %s\nData: %s", err, data) //TODO 這個log應該要統一處理
	}

	return err
}

// GameServer Monitor
func serverRigister() {
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	fmt.Println("[GameServer Monitor]: Start at", now.In(local))

	//初始化
	addServerChan = make(chan GameServer)
	reConnServer = make(chan GameServer)
	serverMap = make(map[string]GameServer)
	pushChan = make(chan Push)

	for {
		select {
		case gameServer := <-addServerChan:
			serverMap[gameServer.Token] = gameServer
			fmt.Println("[GameServer Monitor]: Regist server partner_id:", gameServer.PartnerID, "ip:", gameServer.IP, "token:", gameServer.Token)
		case gameServer := <-reConnServer:
			tmpServer := serverMap[gameServer.Token]
			tmpServer.Conn = gameServer.Conn
			tmpServer.TTL = gameServer.TTL

			serverMap[gameServer.Token] = tmpServer
			fmt.Println("[GameServer Monitor]: Reconnection server partner_id:", gameServer.PartnerID, "ip:", gameServer.IP, "token:", gameServer.Token)
		case push := <-pushChan:
			w, err := push.Target.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write([]byte(push.Data))
			if err := w.Close(); err != nil {
				return
			}
			// push.Target.WriteMessage(websocket.TextMessage, []byte(push.Data)) //TODO 觀察後移除
		}
	}
}

// 刪除GameServer逾時連線
func serverTimeout() {
	if len(serverMap) > 0 {
		for token, info := range serverMap {
			select {
			case <-info.TTL.C:
				delete(serverMap, token)
				fmt.Println("[GameServer Monitor]: Timeout partner_id:", info.PartnerID, "ip:", info.IP, "token:", info.Token)
			default:
			}
		}
	}
}

//GetSererMap 取已註冊連線清單
func GetSererMap() map[string]GameServer {
	return serverMap
}
