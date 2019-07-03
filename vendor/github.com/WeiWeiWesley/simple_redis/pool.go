package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//P connection pool
type P struct {
	Pool *redis.Pool
}

var redisConn = make(map[string]*P)

//SetPool set connection pool
func SetPool(hostName, host string, connLimit int) {
	redisConn[hostName] = &P{Pool: newPool(host, connLimit)}
}

//GetPool get connection pool
func GetPool(host string) *P {
	if p, ok := redisConn[host]; ok {
		return p
	}
	return nil
}

//CloseAllPool close all connection in pool
func CloseAllPool() {
	for i := range redisConn {
		redisConn[i].Close()
	}
}

func newPool(server string, connLimit int) *redis.Pool {

	return &redis.Pool{
		Wait:        true,
		MaxIdle:     20,
		MaxActive:   connLimit,
		IdleTimeout: 180 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// AddConnect new connection
func AddConnect(host string) *P {
	// fmt.Println("new connection", host)
	return &P{Pool: newPool(host, 50)}
}