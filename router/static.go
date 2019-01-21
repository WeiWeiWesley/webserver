package router

import "github.com/gin-gonic/gin"

//StartStaticRouter StartStaticRouter
func StartStaticRouter() {
	r := gin.Default()

	r.Static("/pic", "./public/pic")
	r.Static("/burger", "./public/burger")

	r.Run(":3900")
}
