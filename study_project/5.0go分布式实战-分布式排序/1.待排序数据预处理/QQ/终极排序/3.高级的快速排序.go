package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mylist := makeArr()
	//mylist := []int{4, 3, 8, 1, 6, 5, 7, 9, 2, 10, 11}
	fmt.Println("start", mylist)
	mylist = QuickSortCallAD(mylist)
	fmt.Println("end", mylist)
}

func makeArr() []int { //造数组
	var length = 101
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)
	return list
}

func QuickSortCallAD(arr []int) []int {
	if len(arr) < 100 { //数组个数小于100时，二分插入排序大于快速排序
		return BinSearchInsertSort(arr)
	} else {
		QuickSortFinal(arr, 0, len(arr)-1)
		return arr
	}
}

func BinSearchInsertSort(mylist []int) []int {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindIndexMid(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                          //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}
}

//二分查找
func FindIndexMid(list []int, start int, end int, cur int) int {
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
		return FindIndexMid(list, start, mid, cur)
	} else {
		return FindIndexMid(list, mid+1, end, cur)
	}
}

//快速排序，递归
func QuickSortFinal(arr []int, left int, right int) {
	if right-left < 100 { //调用插入排序小于100插入快 递归调用可能多次使用
		BinSearchInsertSortIndex(arr, left, right) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		SwapData(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr[left]                                //备份中间值
		It := left                                        // 左段arr[left+1,lt]  <vdata  lt++ 左边第二个值到it++的值小于vdata
		gt := right + 1                                   //右段arr[gt...right]>vdata   gt--
		i := left + 1                                     // 中间段arr[lt+1, i] ==vdata   i++

		//	假如arr=4 7 8 9  4 1 2  3
		//  i=1 vdata=arr[[left]]=4
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
				SwapData(arr, i, It+1) //移动小于的地方，从left的下一个元素开始，i=It+1时可以不变
				It++
				i++
			} else if arr[i] < vdata { //吧最右边大于4的数字与最左边小于4的数交换
				SwapData(arr, i, gt-1)
				gt--
			} else {
				i++ //相等
			}
		}
		SwapData(arr, left, It)         //相当于是left的值是  左段最大的 所以调换到后面 接下来递归
		QuickSortFinal(arr, left, It-1) //递归处理左边
		QuickSortFinal(arr, gt, right)  //递归处理右边
	}
}

//对制定数据段排序
func BinSearchInsertSortIndex(mylist []int, start int, end int) []int {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindIndexMid(mylist, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                              //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//数据交换 不用直接创建新数据节省空间
func SwapData(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
