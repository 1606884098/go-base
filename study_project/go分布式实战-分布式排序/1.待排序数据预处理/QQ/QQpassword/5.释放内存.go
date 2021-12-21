package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {
	arr := make([]int, 0)
	for i := 0; i < 100000000; i++ {
		arr = append(arr, i)
	}
	time.Sleep(time.Second * 10)
	fmt.Println(len(arr))
	arr = nil
	runtime.GC() //建议系统自动回收内存
	debug.FreeOSMemory()
	fmt.Println("内存回收了")
	time.Sleep(time.Second * 100)

}
