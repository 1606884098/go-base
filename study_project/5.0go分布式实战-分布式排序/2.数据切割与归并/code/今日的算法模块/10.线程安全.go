package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func add(paddr *int) {
	for i := 0; i < 1000000; i++ {
		*paddr++
	}
	wg.Done()
}

//[],map,list

func main() {
	var num int = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add(&num)
	}
	wg.Wait()
	fmt.Println(num)

}
