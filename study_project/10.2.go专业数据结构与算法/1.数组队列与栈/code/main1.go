package main

import (
	"./ArrayList"
	"./CricleQueue"
	"./Queue"
	"./StackArray"
	"fmt"
)

func main1() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}
func main2() {
	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	fmt.Println(list.TheSize) //小写是私有只能内部用，大写公有，
}

func main3() {

	list := ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	fmt.Println(list.TheSize) //小写是私有只能内部用，大写公有，
}

func main4() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	fmt.Println(list) //小写是私有只能内部用，大写公有，
}

func main5() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i := 0; i < 4; i++ {
		list.Insert(1, "x5")
		fmt.Println(list)
	}
	fmt.Println("delete")
	list.Delete(5)    //删除
	fmt.Println(list) //小写是私有只能内部用，大写公有，
}

func main6() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	for i := 0; i < 55; i++ {
		list.Insert(1, "x5")
		fmt.Println(list)
	}

	//fmt.Println(list) //小写是私有只能内部用，大写公有，
}

func main7() {
	//定义接口对象，赋值的对象必须实现接口的所有方法
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	list.Append("d3")
	list.Append("f3")

	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next("111111")
		if item == "d3" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)

}
func main8() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}
func main9() {
	mystack := ArrayList.NewArrayListStackX()
	mystack.Push(1)
	fmt.Println(mystack.Pop())
	mystack.Push(2)
	fmt.Println(mystack.Pop())
	mystack.Push(3)
	fmt.Println(mystack.Pop())
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	mystack.Push(11)
	mystack.Push(22)
	mystack.Push(33)
	mystack.Push(44)

	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Pop())

	for it := mystack.Myit; it.HasNext(); {
		item, _ := it.Next("111111")
		fmt.Println(item)
	}
}

func Add(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Add(num-1)
	}
}

func main10() {
	//fmt.Println(Add(5))
	mystack := StackArray.NewStack()
	mystack.Push(5)
	last := 0 //保存结果
	for !mystack.IsEmpty() {
		data := mystack.Pop() //取出数据
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystack.Push((data.(int) - 1))
		}
	}
	fmt.Println(last)

}

//5
//4
//3
//2
//1
//0
//f（n）=f(n-1)+f(n-2)
//1
//1
//2
//3
//5
//8
//13
func FAB(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}
func main11() {
	//fmt.Println(FAB(7))
	mystack := StackArray.NewStack()
	mystack.Push(7)
	last := 0 //保存结果
	for !mystack.IsEmpty() {
		data := mystack.Pop() //取出数据
		if data == 1 || data == 2 {
			last += 1
		} else {
			mystack.Push((data.(int) - 2))
			mystack.Push((data.(int) - 1))
		}
	}
	fmt.Println(last)
}

func main12() {
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	myq.EnQueue(14)
	myq.EnQueue(114)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	myq.EnQueue(11114)
	fmt.Println(myq.DeQueue())
}

func main() {

	var myq CricleQueue.CricleQueue
	CricleQueue.InitQueue(&myq)
	CricleQueue.EnQueue(&myq, 1)
	CricleQueue.EnQueue(&myq, 2)
	CricleQueue.EnQueue(&myq, 3)
	CricleQueue.EnQueue(&myq, 4)
	CricleQueue.EnQueue(&myq, 5)
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
}
