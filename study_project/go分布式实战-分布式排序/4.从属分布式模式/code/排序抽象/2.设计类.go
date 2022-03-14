package main

import "fmt"
import "sync"
import "math/rand"

type QuickSortData struct {
	Data         []interface{}
	IsSmalltoBig bool                                                   //要么从大大小，要么从小到大
	myfunc       func(data1, data2 interface{}, IsSmalltoBig bool) bool //函数指针 目的是用于比较大小
}

func (qdata *QuickSortData) QuickSort() {
	fmt.Println("原始数据", qdata.Data)
	if len(qdata.Data) < 1 {
		//插入排序
		//fmt.Println("插入排序")
		qdata.Data = BinSearchSortI(qdata.Data, qdata.myfunc, qdata.IsSmalltoBig)
	} else {
		//快速排序
		fmt.Println("快速排序")
		QuickSortI(qdata.Data, 0, len(qdata.Data)-1, qdata.myfunc, qdata.IsSmalltoBig)
	}
	fmt.Println("最终数据", qdata.Data)
}

func FindindexMidI(list []interface{}, start int, end int, cur int, myfunc func(data1, data2 interface{}, IsSmalltoBig bool) bool, IsSmalltoBig bool) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		//if list[start]>list[cur]{
		if myfunc(list[start], list[cur], !IsSmalltoBig) {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	//if list[mid]<list[cur]{
	if myfunc(list[mid], list[cur], IsSmalltoBig) {
		return FindindexMidI(list, start, mid, cur, myfunc, IsSmalltoBig)
	} else {
		return FindindexMidI(list, mid+1, end, cur, myfunc, IsSmalltoBig)
	}

}
func BinSearchSortI(mylist []interface{}, myfunc func(data1, data2 interface{}, IsSmalltoBig bool) bool, IsSmalltoBig bool) []interface{} {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindindexMidI(mylist, 0, i-1, i, myfunc, IsSmalltoBig) //0,0,  0,1,  0,2,   0,3
			fmt.Println(p)
			if p != i { //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//对制定数据段排序
func BinSearchSortIndexI(mylist []interface{}, start int, end int, myfunc func(data1, data2 interface{}, IsSmalltoBig bool) bool, IsSmalltoBig bool) []interface{} {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMidI(mylist, start, i-1, i, myfunc, IsSmalltoBig) //0,0,  0,1,  0,2,   0,3
			if p != i {                                                     //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//4 123 697
//123  4  697

//数据交换
func SwapI(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//快速排序，递归
func QuickSortI(arr []interface{}, left int, right int, myfunc func(data1, data2 interface{}, IsSmalltoBig bool) bool, IsSmalltoBig bool) {
	if right-left < 1 {
		BinSearchSortIndexI(arr, left, right, myfunc, IsSmalltoBig) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		SwapI(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr[left]                             //备份中间值
		It := left                                     // arr[left+1,lt]  <vdata  lt++
		gt := right + 1                                //arr[gt...right]>vdata   gt--
		i := left + 1                                  // arr[lt+1, i] ==vdata   i++

		//	4 7 8 9  4 1 2  3
		//  i=1 vdata=4
		//	4 3 8 9  4 1 2 7
		//  i=1 vdata=4
		//	4 3 8 9  4 1 2 7
		//  i=1 vdata=4
		//	4 3 2 9  4 1 8 7
		//  i=2 vdata=4
		//	4 3 2 1  4   9 8 7
		//  i=2 vdata=4
		//1 3 2     4 4      9 8 7

		//	4 7 8 9  4  5 6
		//  //	44    7 8 9   5 6

		for i < gt { //循环到重合
			if myfunc(arr[i], vdata, IsSmalltoBig) {
				SwapI(arr, i, It+1) //移动小于的地方
				It++
				i++

			} else if myfunc(arr[i], vdata, !IsSmalltoBig) { //吧最右边大于4的数字与最左边小于4的数交换
				SwapI(arr, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		SwapI(arr, left, It)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			QuickSortI(arr, left, It-1, myfunc, IsSmalltoBig) //递归处理左边
			wg.Done()
		}()
		go func() {
			QuickSortI(arr, gt, right, myfunc, IsSmalltoBig) //递归处理右边
			wg.Done()
		}()
		wg.Wait()

	}
}

func main() {
	mydata := new(QuickSortData)
	mydata.Data = []interface{}{1, 9, 2, 8, 3, 7, 4, 6, 5}
	mydata.IsSmalltoBig = true
	mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
		if IsSmalltoBig {
			return data1.(int) < data2.(int)
		} else {
			return data1.(int) > data2.(int)
		}
	}
	mydata.QuickSort()
}
func main2() {
	mydata := new(QuickSortData)
	mydata.Data = []interface{}{"x", "a", "z", "b"}
	mydata.IsSmalltoBig = true
	mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
		if IsSmalltoBig {
			return data1.(string) < data2.(string)
		} else {
			return data1.(string) > data2.(string)
		}
	}
	mydata.QuickSort()
}

func main3() {
	type QQ struct {
		password string
		time     int
	}
	mydata := new(QuickSortData)
	mydata.Data = []interface{}{QQ{"asdsa", 1}, QQ{"csdsa", 3}, QQ{"bsdsa", 2}}
	mydata.IsSmalltoBig = true
	mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
		if IsSmalltoBig {
			return data1.(QQ).time < data2.(QQ).time
		} else {
			return data1.(QQ).time > data2.(QQ).time
		}
	}
	mydata.QuickSort()
}
func main23() {
	type QQ struct {
		password string
		time     int
	}
	mydata := new(QuickSortData)
	mydata.Data = []interface{}{QQ{"asdsa", 1}, QQ{"csdsa", 3}, QQ{"bsdsa", 2}}
	mydata.IsSmalltoBig = true
	mydata.myfunc = func(data1, data2 interface{}, IsSmalltoBig bool) bool {
		if IsSmalltoBig {
			return data1.(QQ).password < data2.(QQ).password
		} else {
			return data1.(QQ).password > data2.(QQ).password
		}
	}
	mydata.QuickSort()
}
