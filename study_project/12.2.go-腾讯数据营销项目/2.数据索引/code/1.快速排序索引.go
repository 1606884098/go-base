package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func FindindexMid(list1, list2 []int, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list1[list2[start]] > list1[list2[cur]] {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list1[list2[mid]] < list1[list2[cur]] {
		return FindindexMid(list1, list2, start, mid, cur)
	} else {
		return FindindexMid(list1, list2, mid+1, end, cur)
	}

}
func BinSearchSort(mylist1, mylist2 []int) []int {
	if len(mylist2) <= 1 {
		return mylist2
	} else {
		for i := 1; i < len(mylist2); i++ {
			p := FindindexMid(mylist1, mylist2, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                    //不等，插入
				for j := i; j > p; j-- {
					mylist2[j], mylist2[j-1] = mylist2[j-1], mylist2[j] //数据移动
				}
			}
		}
		return mylist2
	}

}

//对制定数据段排序
func BinSearchSortIndex(mylist1, mylist2 []int, start int, end int) []int {
	if end-start <= 1 {
		return mylist2
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMid(mylist1, mylist2, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                        //不等，插入
				for j := i; j > p; j-- {
					mylist2[j], mylist2[j-1] = mylist2[j-1], mylist2[j] //数据移动
				}
			}
		}
		return mylist2
	}

}
func QuickSortCall(arr1, arr2 []int) []int {
	if len(arr2) <= 5 {
		return BinSearchSort(arr1, arr2)
	} else {
		QuickSort(arr1, arr2, 0, len(arr1)-1)
		return arr2
	}
}

//4 123 697
//123  4  697

//数据交换
func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//快速排序，递归
func QuickSort(arr1, arr2 []int, left int, right int) {
	if right-left < 5 {
		BinSearchSortIndex(arr1, arr2, left, right) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		Swap(arr2, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr2[left]                            //备份中间值
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
			if arr1[arr2[i]] > arr1[vdata] {
				Swap(arr2, i, It+1) //移动小于的地方
				It++
				i++

			} else if arr1[arr2[i]] < arr1[vdata] { //吧最右边大于4的数字与最左边小于4的数交换
				Swap(arr2, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		Swap(arr2, left, It)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			QuickSort(arr1, arr2, left, It-1) //递归处理左边
			wg.Done()
		}()
		go func() {
			QuickSort(arr1, arr2, gt, right) //递归处理右边
			wg.Done()
		}()
		wg.Wait()

	}
}

func main() {

	mylist1 := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	mylist2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("start", mylist1)
	mylist2 = QuickSortCall(mylist1, mylist2)
	fmt.Println("end", mylist1)
	fmt.Println("end", mylist2)
	for i := 0; i < len(mylist2); i++ {
		fmt.Println(mylist1[mylist2[i]])
	}

}

func main1z() {

	mylist1 := []int{1, 9, 2, 8, 3}
	mylist2 := []int{0, 1, 2, 3, 4}
	//mylist2:=[]int{1,3,4,2,0}
	fmt.Println("end", mylist1)
	fmt.Println("end", mylist2)
	for i := 0; i < len(mylist2); i++ {
		fmt.Println(mylist1[mylist2[i]])
	}

}

//interface 需要两个Interface
//字符串对应字符串
//数字对应字符串
//字串对应结构体
//数字对应数字
