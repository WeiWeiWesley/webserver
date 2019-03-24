package main

import (
	"webserver/router"
	"webserver/router/ws"
)

func init() {

}

func main() {

	go router.StartAPIRouter() //port 3700
	go router.StartStaticRouter() //port 3800
	go ws.StartWsRouter() //port 3900

	select {}
}
