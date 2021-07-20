package main

import "fmt"

func main() {
	//遍历channel
	forEachChannel()
}

func forEachChannel() {
	forEachChan := make(chan interface{}, 200)
	for i := 0; i < 100; i++ {
		forEachChan <- i
		forEachChan <- "cc"
	}
	a := <-forEachChan
	fmt.Println(a)
	close(forEachChan)           //遍历chan需要关闭，否则会报错deallock错误
	for v := range forEachChan { //遍历管道只能用for range 不能用普通的for遍历
		fmt.Println("v=", v)
	}
}
