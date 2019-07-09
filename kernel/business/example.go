package business

import (
	"net/http"
	"webserver/kernel/common"
	"webserver/kernel/db/mysql"
	"webserver/kernel/redis"

	"github.com/gin-gonic/gin"
)

//GetRedisData GetRedisData
func GetRedisData(c *gin.Context) {
	conn := redis.GetPool(common.RedisPoolKey)

	res, err := conn.KEYS(`*`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})

		return
	}

	c.JSON(http.StatusOK, common.HTTPResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    res,
	})
}

//PingMySQL PingMySQL
func PingMySQL(c *gin.Context) {
	db, err := mysql.GetPool(common.MySQLPoolKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})

		return
	}

	pingErr := db.Conn.DB().Ping()
	if pingErr != nil {
		c.JSON(http.StatusInternalServerError, common.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})

		return
	}

	c.JSON(http.StatusOK, common.HTTPResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    "ping success",
	})
}
