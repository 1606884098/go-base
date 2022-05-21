package main

import (
	"fmt"
	"sync"
)

type Single struct {
	data int
}

var singleton *Single
var once sync.Once //内核信号，时时刻刻智能运行一个

func GetInterface() *Single {
	//once.Do(func (){singleton=&Single{100}}) 单例
	singleton = &Single{100}
	return singleton
}
func main() {
	i1 := GetInterface()
	i2 := GetInterface()
	if i1 == i2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不等")
	}
}
