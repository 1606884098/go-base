package main

import "math/rand"
import "time"
import "fmt"
import "sync"

func makeArrConcurrent() []int {
	var length = 30
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	//fmt.Println(list)
	return list
}

//123  4  56789

func QuickGoSort(data []int) []int {
	if len(data) <= 1 {
		return data
	} else {
		var wg sync.WaitGroup //批量等待 法令枪
		c := data[0]          //第一个数
		var left, mid, right []int
		mid = append(mid, c) //第一个数
		for k, v := range data {
			if k == 0 {
				continue
			}
			if c < v {
				right = append(right, v)
			} else if c > v {
				left = append(left, v)
			} else {
				mid = append(mid, v)
			}
		}

		go func() {
			left = QuickGoSort(left)
			wg.Done()
		}()
		go func() {
			right = QuickGoSort(right)
			wg.Done()
		}()
		wg.Add(2)

		wg.Wait() //等待两个线程返回

		data := []int{}
		if len(left) > 0 {
			data = append(data, left...)
		}
		data = append(data, mid...)
		if len(right) > 0 {
			data = append(data, right...)
		}
		return data

	}

}

func main() {
	data := makeArrConcurrent()
	fmt.Println(data)
	fmt.Println(QuickGoSort(data))
}
