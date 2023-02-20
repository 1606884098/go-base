package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
select 一定程度上可以类比于 linux 中的 IO 多路复用中的 select。后者相当于提供了对多个 IO
事件的统一管理，而 Golang 中的 select 相当于提供了对多个 channel 的统一管理。当然这只是
select 在 channel 上的一种使用方法。

值得注意的是 select 中的 break 只能跳到 select 这一层。select 使用的时候一般配合 for
循环使用，像下面这样，因为正常 select 里面的流程也就执行一遍。这么来看 select 中的 break
就稍显鸡肋了。所以使用 break 的时候一般配置 label 使用，label 定义在 for 循环这一层。
*/
type a struct {
}

func main() {
	var q struct{}
	ch1 := make(chan int, 10)
	stopCh := make(chan struct{}, 1)
	stopCh <- q
	ch1 <- 11
	ch1 <- 12
	for {
		return
		/*select {
		case <-stopCh:
			fmt.Println("返回了")
			return
		case <-ch1:
			fmt.Println("liwenzhou.com")

		}*/
	}

	fmt.Println("也有执行")
	ch := make(chan int, 1)
	ch <- 1
	ch2 := make(chan int, 10)
	ch2 <- 13
	ch2 <- 14
	//for{//顺序取值,不能这么写，必须知道取多少次
	for i := 1; i <= 2; i++ {
		// 尝试从ch1接收值
		d1, ok := <-ch1
		if ok {
			fmt.Println(d1)
		}
		// 尝试从ch2接收值
		d2, ok := <-ch2
		if ok {
			fmt.Println(d2)
		}
	}
	for {
		select { //随机取
		case ch2 <- 10: //写入数据
			fmt.Println("取：", len(ch2))
		case x := <-ch2:
			fmt.Println("写：", x) //取出数据发送到第三方
		default: //如果上面两个case都取完写满了会执行下面的语句
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("没操作")
		}
	}

}
