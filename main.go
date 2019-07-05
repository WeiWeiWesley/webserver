package main

import (
	"os"

	"webserver/router"
	"webserver/router/ws"
)

func init() {
	//Check ENV
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "local")
	}
}

func main() {
	go router.StartAPIRouter()    //port 3700
	go router.StartStaticRouter() //port 3800
	go ws.StartWsRouter()         //port 3900

	select {}
}
