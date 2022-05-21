package main

import (
	"fmt"
	"time"
)

//  5 1  3  2  4
//  5 1  3   2 4
//写入 1 us
//1亿数据
//多线程，分布式
var flag bool
var container chan bool
var count int

func main() {

	var array []int = []int{16, 8, 1, 24, 30}
	flag = true                    //标识，区分
	container = make(chan bool, 5) //5个管道
	for i := 0; i < len(array); i++ {
		go tosleep(array[i])
	}
	go listen(len(array))
	for flag {
		time.Sleep(1 * time.Second)
	}

}

func listen(size int) {
	for flag {
		select {
		case <-container:
			count++            //计数器
			if count >= size { //等待5个数字采集完成就退出
				flag = false
				break
			}
		}
	}
}

func tosleep(data int) {
	time.Sleep(time.Duration(data) * time.Microsecond * 1000)
	fmt.Println("sleep", data)
	container <- true //管道输入ok
}
