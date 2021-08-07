package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲连接
		MaxActive:   maxActive,   //数据库的最大连接
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接
			return redis.Dial("tcp", address)
		},
	}
}
