package main

import (
	"context"
	"fmt"
	//"github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis/v8"
)

//定义一个全局的 pool
/*var pool *redis.Pool


//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个 ip 的 redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}*/

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		//	端口需要改，这里是docker的端口
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           6,  // 指定库 DB
		PoolSize:     15,
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
		//超时
		//DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		//ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		//WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		//PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	})
}

func main() {
	//先从 pool 取出一个链接
	/*conn := pool.Get()
	defer conn.Close()
	//可以插入数据到指定库
	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return

	}

	fmt.Println("r=", r)

	//如果我们要从 pool 取出链接，一定保证链接池是没有关闭
	//pool.Close() conn2 := pool.Get()

	_, err = conn.Do("Set", "name2", "汤姆猫~~2")
	if err != nil {
		fmt.Println("conn.Do err~~~~=", err)
		return
	}

	//取出
	r2, err := redis.String(conn.Do("Get", "name2"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r=", r2)

	//fmt.Println("conn2=", conn2)*/

	//	添加key
	//0表示没有过期时间
	rdb.Set(ctx, "testKey121123", "xxx", 0)

	//	获取值
	val, err := rdb.Get(ctx, "testKey").Result()
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("值：", val)

}
