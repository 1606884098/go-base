package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type bird struct {
	fly  int
	name string
}
type birdHeap []bird

func (h *birdHeap) Push(x interface{}) {
	panic("implement me")
}

//比大小
func (h *birdHeap) Less(i, j int) bool {
	return (*h)[i].fly < (*h)[j].fly
}

func (h *birdHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *birdHeap) Len() int {
	return len(*h)
}

func (h *birdHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *birdHeap) Remove(idx int) interface{} {
	h.Swap(idx, h.Len()-1)
	return h.Pop()
}

//先将数据放到切片里面，这是初始数据，然后    heap.Init(h) 就是将数据利用堆结构排列，最后将数据调用 方法来操作
func main() {
	h := new(birdHeap)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(h)
	for i := 0; i < 10; i++ {

	}
	heap.Init(h)
}
