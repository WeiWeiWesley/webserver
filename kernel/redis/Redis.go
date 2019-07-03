package redis

import (
	"github.com/WeiWeiWesley/simple_redis"
)

//NewPool Add a new pool
func NewPool(hostName, host string, connLimit int) {
	redis.SetPool(hostName, host, connLimit)
}

//GetPool Fetch a conn pool by name
func GetPool(poolName string) *redis.P {
	return redis.GetPool(poolName)
}