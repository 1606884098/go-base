package main

import "container/list"
import "fmt"

func MergeStack(arr []string) string {
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
			v3 := v1 + v2 //可以是两个文件或者数组
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
	arrlist := []string{"A", "B", "C", "D", "E", "F"} //这里相当于是6个文件名
	fmt.Println("last", MergeStack(arrlist))
}
