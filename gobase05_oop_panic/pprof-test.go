package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	// 开启pprof，监听请求

	//开启对锁调用的跟踪
	runtime.SetMutexProfileFraction(1)
	//开启对阻塞操作的跟踪 这两个不开启捕捉不到的
	runtime.SetBlockProfileRate(1)

	ip := "localhost:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}

}
