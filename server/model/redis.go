package model

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxActive:   maxActive,   //数据库最大连接数，0表示没有限制
		MaxIdle:     maxIdle,     //最大空闲连接数
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接的代码，链接哪个api的redis
			return redis.Dial("tcp", address)
		},
	}
}
