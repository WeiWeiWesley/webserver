package router

import "github.com/gin-gonic/gin"

//StartAPIRouter StartAPIRouter
func StartAPIRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3700") // listen and serve on 0.0.0.0:8080
}
