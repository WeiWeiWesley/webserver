package main

import (
	"os"

	"webserver/kernel/common"
	"webserver/kernel/redis"
	"webserver/router"
	"webserver/router/ws"
)

func init() {
	//Check ENV
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "local")
	}
	
	redis.NewPool(common.RedisPoolKey, "127.0.0.1:6379", 50)
}

func main() {
	go router.StartAPIRouter()    //port 3700
	go router.StartStaticRouter() //port 3800
	go ws.StartWsRouter()         //port 3900

	select {}
}
