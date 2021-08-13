package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist := list.New() //切片是连续的存储空间查找快，链表增删快  比如windows的任务管理器就是链表实现的
	for i := 0; i < 100; i++ {
		mylist.PushFront(i)
	}
	for it := mylist.Front(); it != nil; it.Next() {
		fmt.Println(it.Value, it.Prev(), it.Next())
	}
}
