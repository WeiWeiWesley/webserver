package main

import (
	"webserver/router"
)

func init() {

}

func main() {

	go router.StartAPIRouter()

	go router.StartStaticRouter()

	select {}
}
