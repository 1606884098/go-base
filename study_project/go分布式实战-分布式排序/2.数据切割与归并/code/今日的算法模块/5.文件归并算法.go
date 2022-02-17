package main

import "container/list"
import "fmt"

func Merge(arr []string) string {
	mylist := list.New()
	for i := 0; i < len(arr); i++ {
		mylist.PushBack(arr[i]) //数据批量压入
	}
	fmt.Println(mylist.Len()) //栈数据长度

	for mylist.Len() != 1 {
		e1 := mylist.Back()
		mylist.Remove(e1)

		e2 := mylist.Back()
		mylist.Remove(e2) //取得两个数据

		if e1 != nil && e2 != nil { //两个数据不为空，归并
			v1, _ := e1.Value.(string)
			v2, _ := e2.Value.(string)
			v3 := v1 + v2
			mylist.PushBack(v3)
		} else if e1 != nil && e2 == nil { //一个不为空，另外一个为空，再次压入
			v1, _ := e1.Value.(string)
			mylist.PushBack(v1)
		} else if e1 == nil && e2 == nil { //均为空，跳出循环
			break
		} else {
			break
		}

	}

	return mylist.Back().Value.(string)

}

func main() {
	arrlist := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println("last", Merge(arrlist))
}

func main1() {
	mylist := list.New()
	for i := 0; i < 10; i++ {
		mylist.PushBack(i)
	}
	for mylist.Len() != 0 {
		e := mylist.Back()
		mylist.Remove(e)
		fmt.Println(e.Value.(int))
	}

}

// 1  2  3  4  5  6  7
// 12  34  56  7
//  1234  567
//  1234567
