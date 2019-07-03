package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//GET key
func (p *P) GET(key string) ([]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bytes(conn.Do("GET", key))
}

//SET key value
func (p *P) SET(key string, value string) error {
	conn := p.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value); if err != nil {
		if len(value) > 15 {
			value = value[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, value, err)
	}

	return err
}

//EXISTS key [key ...]
func (p *P) EXISTS(key string) (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("EXISTS", key))
}

//DEL DEL key [key ...]
func (p *P) DEL(key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Int64(conn.Do("DEL", key))
}

//GetKeys get key by pattern
func (p *P) GetKeys(pattern string) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern)); if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

//INCR key
func (p *P) INCR(counterKey string) (int, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}

//HGET key field
func (p *P) HGET(key string, field string) ([]byte, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bytes(conn.Do("HGET", key, field))
}

//HGETALL Returns all fields and values of the hash stored at key.
func (p *P) HGETALL(key string) (map[string]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.StringMap(conn.Do("HGETALL", key))
}

//HSET HSET key field value
func (p *P) HSET(key string, field string, value string) (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("HSET", key, field, value))
}

//HEXISTS HEXISTS key field
func (p *P) HEXISTS(key string, field string) (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("HEXISTS", key, field))
}

//HDEL HDEL key field [field ...]
func (p *P) HDEL(key string, field string) (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	 return redis.Bool(conn.Do("HDEL", key, field))
}

//HMSET key field value [field value ...]
func (p *P) HMSET(key string, hashData map[interface{}]interface{}) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(hashData)...)
}

//KEYS KEYS pattern
func (p *P) KEYS(pattern string) ([]string, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Strings(conn.Do("KEYS", pattern))
}

//FLUSHALL [ASYNC]
func (p *P) FLUSHALL() (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()


	return redis.Bool(conn.Do("FLUSHALL"))
}

//Close close connection
func (p *P) Close() error {
	err := p.Pool.Close()

	return err
}

//EXPIRE EXPIRE key seconds
func (p *P) EXPIRE(key string, sec int) (interface{}, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return conn.Do("EXPIRE", key, sec)
}

//HINCRBYFLOAT HINCRBYFLOAT key field increment
func (p *P) HINCRBYFLOAT(key, field string, incr float64) (float64, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Float64(conn.Do("HINCRBYFLOAT", key, field, incr))
}

//TTL Returns the remaining time to live of a key that has a timeout.
func (p *P) TTL(key string) (int64, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Int64(conn.Do("TTL", key))
}

//LPUSH Insert all the specified values at the head of the list stored at key.
func (p *P) LPUSH(key, value string) (bool, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("LPUSH", key, value))
}

//RPOP Removes and returns the last element of the list stored at key.
func (p *P) RPOP(key string) (string, error) {
	conn := p.Pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("RPOP", key))
}