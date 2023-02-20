package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	mutex = sync.Mutex{} //锁是成员变量，不能定义为局部变量，确保同一把锁

	c = make(chan int)
)

/*
	var 变量 chan 元素类型(要传输的类型)
	1.1.3. 创建channel
	通道是引用类型，通道类型的空值是nil。
	var ch chan int
	fmt.Println(ch) // <nil>
	声明的通道后需要使用make函数初始化之后才能使用。
	创建channel的格式如下：
	make(chan 元素类型, [缓冲大小])
	channel的缓冲大小是可选的。
	1.1.4. channel操作
	通道有发送（send）、接收(receive）和关闭（close）三种操作。
	发送和接收都使用<-符号。
	现在我们先使用以下语句定义一个通道：
	ch := make(chan int)
	发送
	将一个值发送到通道中。
	ch <- 10 // 把10发送到ch中
	接收
	从一个通道中接收值。
	x := <- ch // 从ch中接收值并赋值给变量x
	<-ch       // 从ch中接收值，忽略结果
	关闭
	我们通过调用内置的close函数来关闭通道。
	close(ch)
	关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关
	闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必
	须要做的，但关闭通道不是必须的。关闭后的通道有以下特点：
	1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。

	无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：
	func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("发送成功")
	}
	上面这段代码能够通过编译，但是执行的时候会出现以下错误：
	    fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [chan send]:
    main.main()
            .../src/github.com/pprof/studygo/day06/channel02/main.go:8 +0x54
	为什么会出现deadlock错误呢？
	因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候
	才能发送值。

	无缓冲的通道必须有接收才能发送。
	上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？
	一种方法是启用一个goroutine去接收值，例如：
	func recv(c chan int) {
    ret := <-c
    fmt.Println("接收成功", ret)
	}
	func main() {
    ch := make(chan int)
    go recv(ch) // 启用goroutine从通道接收值
    ch <- 10
    fmt.Println("发送成功")
	}
	无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成
	功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一
	个goroutine在该通道上发送一个值。使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。
	因此，无缓冲通道也被称为同步通道。

	我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：
	func main() {
    ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
    ch <- 10
    fmt.Println("发送成功")
	}
	只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	1.1.7. close()
	可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
	我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往
	该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭
	了呢？
	我们来看下面这个例子：
	// channel 练习
	func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    // 开启goroutine将0~100的数发送到ch1中
    go func() {
        for i := 0; i < 100; i++ {
            ch1 <- i
        }
        close(ch1)
    }()
    // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
    go func() {
        for {
            i, ok := <-ch1 // 通道关闭后再取值ok=false
            if !ok {
                break
            }
            ch2 <- i * i
        }
        close(ch2)
    }()
    // 在主goroutine中从ch2中接收值打印
    for i := range ch2 { // 通道关闭后会退出for range循环
        fmt.Println(i)
    }
	}
	从上面的例子中我们看到有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。
	单向chan
 	1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
    2.<-chan int是一个只能接收的通道，可以接收但是不能发送。

	1) channel 中只能存放指定的数据类型
	2) channle 的数据放满后，就不能再放入了
	3) 如果从 channel 取出数据后，可以继续放入
	4) 在没有使用协程的情况下，如果 channel 数据取完了，再取，就会报 dead lock


*/
func main1() {
	//遍历channel
	//forEachChannel()

	// 我们这里开启多个协程完成这个任务[200 个] map版
	/*for i := 1; i <=30; i++ {
		go test(i)
	}

	time.Sleep(time.Second) //等待上面的协程执行完毕
	for k, v := range myMap {
		fmt.Printf("%d ----> %s\n", k, v)
	}*/

	//time.Sleep(time.Second) //等待上面的协程执行完毕
	start := time.Now()
	go receive()
	//go receive()//开启多个速度更慢
	// 我们这里开启多个协程完成这个任务[200 个]channel版
	for i := 1; i <= 500; i++ {
		//go test1(i)//会导致数据丢失
		//test1(i)
		c <- i
	}
	//close(c)
	//go receive() 放这里报错
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

func receive() {
	i := 1
	/*	select {
		case cha:=<-c:
			println("c的值：", cha)
		default:
			println("c的值：无")
		}*/
	for {
		cha := <-c
		//myMap[i]=c//并发
		i++
		println("c的值：", cha)
	}

}

func test1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	c <- res
}

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func test(n int) { //性能差
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mutex.Lock()   //不加锁会并发报错
	myMap[n] = res //concurrent map writes?
	mutex.Unlock()
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
