package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println("只执行一次")
		})
	}
	var ch chan int
	ch = make(chan int, 1)
	fmt.Println(ch)
	ch <- 10
	//ch <-10
	//ch <-10
	/*	go func() {
		ch <-10
		ch <-10
	}()*/
	//ch<-10

	//ch <-10
	//fmt.Println(<-ch)
	//close(ch)
	/*dataCh1 := make(chan int)
	go func() {

			println(<-dataCh1)

	}()
	dataCh1<-1
	dataCh1<-2
	dataCh1<-3
	dataCh1<-4
	dataCh1<-5
	close(dataCh1)
	println(<-dataCh1)
	println(<-dataCh1)
	println(<-dataCh1)
	println(<-dataCh1)
	a,ok:=<-dataCh1
	println(ok)
	println(a)
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <- time.After(time.Hour):
	}*/
}
