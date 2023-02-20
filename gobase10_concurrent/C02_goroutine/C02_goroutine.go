package main

import (
	"fmt"
	"time"
)

//不要通过共享内存来通信，而要通过通信来实现内存共享
/*
主goroutine退出后，其它的工作goroutine也会自动退出：
*/
func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {
	// 并发执行程序
	go running() //1、普通创建
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)

	go func() { //2、匿名方式创建
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
	var input1 string
	fmt.Scanln(&input1)

}
