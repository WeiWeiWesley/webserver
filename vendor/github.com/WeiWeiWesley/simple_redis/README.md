Simple Redis
===

# Usage

1. init
```golang
import "github.com/WeiWeiWesley/simple_redis"

func init() {
    redis.SetPool(${PoolKey}, ${Host}, ${poolLimit})
    defer redis.CloseAllPool()
}
```

2. use connection pool

```golang
import "github.com/WeiWeiWesley/simple_redis"

func myfunc() {
    Conn := redis.GetPool(${PoolKey})
    res, err := Conn.GET("just key")
    if err != nil {
        //do somthing
    }
}
```