package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}
func entry() {
	var rd Road = Road(999)
	r := &rd
	runtime.SetFinalizer(r, findRoad)
}

/*
尽量使用栈上分配，尽量避免栈上分配变量发生逃逸
go run -gcflags "-m -l" 文件名.go
*/

func main() {
	entry()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
