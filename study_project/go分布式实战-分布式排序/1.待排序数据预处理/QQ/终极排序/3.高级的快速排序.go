package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FindindexMid12(list []int, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list[start] > list[cur] {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list[mid] < list[cur] {
		return FindindexMid12(list, start, mid, cur)
	} else {
		return FindindexMid12(list, mid+1, end, cur)
	}

}
func BinSearchSort1(mylist []int) []int {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindindexMid12(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                            //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//对制定数据段排序
func BinSearchSortIndex1(mylist []int, start int, end int) []int {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMid12(mylist, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}
func QuickSortCall1(arr []int) []int {
	if len(arr) < 10 {
		return BinSearchSort1(arr)
	} else {
		QuickSort1(arr, 0, len(arr)-1)
		return arr
	}
}

//4 123 697
//123  4  697

//数据交换
func Swap1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//快速排序，递归
func QuickSort1(arr []int, left int, right int) {
	if right-left < 10 {
		BinSearchSortIndex1(arr, left, right) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		Swap1(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
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
			if arr[i] > vdata {
				Swap1(arr, i, It+1) //移动小于的地方
				It++
				i++

			} else if arr[i] < vdata { //吧最右边大于4的数字与最左边小于4的数交换
				Swap1(arr, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		Swap1(arr, left, It)

		QuickSort1(arr, left, It-1) //递归处理左边
		QuickSort1(arr, gt, right)  //递归处理右边

	}
}

func main1() {

	mylist := []int{2, 1, 9, 17, 8, 10}
	mylist = BinSearchSortIndex1(mylist, 3, 5)
	fmt.Println(mylist)

}
func main() {
	//mylist := makearr()
	mylist := []int{4, 3, 8, 1, 6, 5, 7, 9, 2, 10, 11}
	fmt.Println("start", mylist)
	mylist = QuickSortCall1(mylist)
	fmt.Println("end", mylist)
}

func makearr() []int { //造数组
	var length = 100
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)
	return list
}
