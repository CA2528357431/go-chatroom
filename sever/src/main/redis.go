package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool


// 初始化redis池
func Initpool(address string,maxidle int,timeout time.Duration)  {

	pool = &redis.Pool{}
	pool.Dial = func() (redis.Conn, error) {
		return redis.Dial("tcp",address)
	}
	pool.MaxIdle = maxidle
	pool.IdleTimeout = timeout

}
