package main

import (
	"webserver/kernel/common"
	"webserver/kernel/db/mysql"
	"webserver/kernel/redis"
	_ "webserver/kernel/rpc"
	"webserver/router"
	"webserver/router/ws"
)

func init() {
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
