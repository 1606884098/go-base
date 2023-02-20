package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义全局chan，存储票总数
var totalTickets1 chan int
var totalTickets2 chan int
var totalTickets3 chan int
var wg sync.WaitGroup
var ticketsFromDB int = 25689 // 初始化票数量：总票张数
var modulusData int = 10000   //每个channel数量是多少

func init() {
	totalTickets1 := make(chan int, 2)
	totalTickets1 <- modulusData
	totalTickets2 := make(chan int, 2)
	totalTickets2 <- modulusData
	totalTickets3 := make(chan int, 2)
	totalTickets3 <- ticketsFromDB - ticketsFromDB%modulusData*modulusData //模数求余方法

}

func main() {

	wg.Add(5)

	go sell("售票口1")
	go sell("售票口2")
	go sell("售票口3")
	go sell("售票口4")
	go sell("售票口5")

	wg.Wait()

	fmt.Println("main over.")
}

func sell(name string) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	for {
		residue1, ok := <-totalTickets1
		if !ok {
			fmt.Printf("%s: Sold Out 1\n", name)
			break
		}
		residue2, ok := <-totalTickets2
		if !ok {
			fmt.Printf("%s: Sold Out 1\n", name)
			break
		}
		residue3, ok := <-totalTickets3
		if !ok {
			fmt.Printf("%s: Sold Out 1\n", name)
			break
		}
		if residue1 > 0 || residue2 > 0 || residue3 > 0 {
			select {
			case totalTickets1 <- residue1 - 1:
				fmt.Println(name, "chan1售出1张票，余票：", residue1)
			case totalTickets2 <- residue2 - 1:
				fmt.Println(name, "chan2售出1张票，余票：", residue2)
			case totalTickets3 <- residue3 - 1:
				fmt.Println(name, "chan3售出1张票，余票：", residue3)
			}
		} else {
			fmt.Printf("%s: Sold Out 2\n", name)
			close(totalTickets1)
			close(totalTickets2)
			close(totalTickets3)
			break
		}
	}
}
