package main

import (
	"os"

	"webserver/kernel/common"
	"webserver/kernel/db/mysql"
	"webserver/kernel/redis"
	"webserver/router"
	"webserver/router/ws"
)

func init() {
	//Check ENV
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "local")
	}

	//LoadConfig
	confing := common.LoadConfig()

	//Add redis conn pool
	redis.NewPool(common.RedisPoolKey, confing.RedisDefault.Host, confing.RedisDefault.MaxConn)

	//Add MySQL db conn pool
	mysql.NewPool(common.MySQLPoolKey, confing.MySQLDefault)
}

func main() {
	defer redis.CloseRedis()

	go router.StartAPIRouter()    //port 3700
	go router.StartStaticRouter() //port 3800
	go ws.StartWsRouter()         //port 3900

	select {}
}
