package main

import "fmt"

func main() {
	//遍历channel
	for_each_channel()
}

func for_each_channel() {
	for_each_chan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		for_each_chan <- i
	}
	a := <-for_each_chan
	fmt.Println(a)
	close(for_each_chan)           //遍历chan需要关闭，否则会报错deallock错误
	for v := range for_each_chan { //遍历管道只能用for range 不能用普通的for遍历
		fmt.Println("v=", v)
	}
}
