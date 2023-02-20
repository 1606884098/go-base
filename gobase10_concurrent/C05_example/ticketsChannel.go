package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 定义全局chan，存储票总数
var totalTickets chan int
var wg sync.WaitGroup
var once sync.Once

func main() {

	/*
		指定用多少线程（cpu的线程4核8线程的线程）运行这个程序,默认最大值运行除非是
		监控，或者日志程序才设置这个参数节省资源运行其他程序
	*/
	runtime.GOMAXPROCS(3) //
	// 初始化票数量：总票数10张
	totalTickets = make(chan int, 10)
	var good int = 10 //从数据库里查出来的数据
	totalTickets <- good
	var goods = make([]int, 5) //从数据库里查出来的数据,动态配置
	for i := 0; i <= len(goods); i++ {
		wg.Add(1) //启动一个协程就加一
		//协程什么时候结束，函数执行完协程也就借宿了.main函数也是一样的
		go sell("售票口" + strconv.Itoa(i))
	}

	/*	wg.Add(5)
		go sell("售票口1")
		go sell("售票口2")
		go sell("售票口3")
		go sell("售票口4")
		go sell("售票口5")*/

	wg.Wait()                 //等待wg的计数器为0
	fmt.Println("main over.") //这样做的目的可以主线程和子协程同步关闭
}

func sell(name string) {
	defer wg.Done() //代表结束当前协程

	rand.Seed(time.Now().UnixNano()) //保证每次执行的时候都有点不一样
	for {
		residue, ok := <-totalTickets
		if !ok {
			fmt.Printf("%s: Sold Out 1\n", name)
			break
		}
		if residue > 0 {
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			totalTickets <- residue - 1
			fmt.Println(name, "售出1张票，余票：", residue)
		} else {
			fmt.Printf("%s: Sold Out 2\n", name)
			once.Do(func() { close(totalTickets) })
			break
		}
	}
}
