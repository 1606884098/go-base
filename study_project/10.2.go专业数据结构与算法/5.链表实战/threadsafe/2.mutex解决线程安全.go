package main

import (
	"fmt"
	"time"
	"sync"
)
//线程安全，多个线程访问同一个资源，产生资源竞争，最终结果不正确
var money int=0
var lock *sync.RWMutex=new(sync.RWMutex) //初始化

func add(pint  *int){
	lock.Lock()
	for i:=0;i<100000;i++{
		*pint++
	}
	lock.Unlock()
}

func main(){
	for i:=0;i<1000;i++{
		go  add(&money)
	}
	time.Sleep(time.Second*20)
	fmt.Println(money)
}

