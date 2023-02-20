package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"time"
)

/*
container/list
*/
func main() {
	mylist := list.New() //切片是连续的存储空间查找快，链表增删快  比如windows的任务管理器就是链表实现的
	for i := 0; i < 10; i++ {
		mylist.PushFront(i)
	}
	for it := mylist.Front(); it != nil; it = it.Next() {
		fmt.Println(it.Value, it.Prev(), it.Next())
	}

	element := mylist.PushBack("尾部添加") //尾部添加返回句柄
	mylist.InsertAfter("在句柄元素后添加", element)
	mylist.InsertBefore("在句柄元素前添加", element)
	mylist.PushFront("头部添加")
	element1 := mylist.PushFront("测试删除")
	mylist.Remove(element1)

	r := ring.New(5) //环形链表
	for i := 0; i < 5; i++ {
		r.Value = i
		r = r.Next()
	}

	testCreat()           //结论是切片创建的性能优于链表
	testfor()             //遍历也就是查询的性能切片优于链表
	testInsertAndRemove() //插入和删除链表的性能优于切片
}

func testInsertAndRemove() {
	slice := make([]int, 10)
	for i := 0; i < 1*100000*1000; i++ {
		slice = append(slice, i)
	}

	list := list.New()
	for i := 0; i < 1*100000*1000; i++ {
		list.PushFront(i)
	}
	t := time.Now()
	sl := slice[:100000*500]
	sf := slice[100000*500:]
	sl = append(sl, 10)
	sf = append(sf, sl...)
	fmt.Println("slice", time.Now().Sub(t).String()) //slice 3.9038ms
}

func testfor() {

	slice := make([]int, 1)
	for i := 0; i < 10000000; i++ {
		slice = append(slice, i)
	}
	start := time.Now()
	for _, _ = range slice {

	}
	fmt.Println("slice", time.Now().Sub(start).String()) //slice 3.9038ms

	list := list.New()
	for i := 0; i < 10000000; i++ {
		list.PushBack(i)
	}
	start1 := time.Now()
	for e := list.Front(); e != nil; e = e.Next() {

	}
	fmt.Println("list", time.Now().Sub(start1).String()) //list 42.9441ms
}

func testCreat() {
	start := time.Now()
	slice := make([]int, 1)
	for i := 0; i < 100000; i++ {
		slice = append(slice, i)
	}
	fmt.Println("slice", time.Now().Sub(start).String()) //slice 3.8842ms

	start1 := time.Now()
	list := list.New()
	for i := 0; i < 100000; i++ {
		list.PushBack(i)
	}
	fmt.Println("list", time.Now().Sub(start1).String()) //list 9.7597ms
}
