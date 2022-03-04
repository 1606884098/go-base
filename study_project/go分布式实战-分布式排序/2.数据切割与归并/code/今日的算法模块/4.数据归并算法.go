package main

import (
	"fmt"
	"math/rand"
	"time"
)

//将两个数组排序，然后再归并到一个数组中
func main() {
	mylist1 := makeArrMerge()
	mylist1 = QuickSortCallMerge(mylist1)
	fmt.Println("end1", mylist1)
	time.Sleep(time.Second)
	mylist2 := makeArrMerge()
	mylist2 = QuickSortCallMerge(mylist2)
	fmt.Println("end2", mylist2)
	fmt.Println(MergeArr(mylist1, mylist2))
}

//归并算法:将两个有序的数组元素比较然后写入到第三个数组中
func MergeArr(arr1 []int, arr2 []int) []int {
	allarr := []int{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) { //i和j有可能不等，超出小的哪个时跳出fof
		if arr1[i] < arr2[j] {
			allarr = append(allarr, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			allarr = append(allarr, arr2[j])
			j++
		} else {
			allarr = append(allarr, arr1[i])
			i++
			allarr = append(allarr, arr2[j])
			j++
		}

	}
	for i < len(arr1) { //i>j的情况，直接将尾部追加
		allarr = append(allarr, arr1[i])
		i++
	}
	for j < len(arr2) { //i<j的情况，直接将尾部追加
		allarr = append(allarr, arr2[j])
		j++
	}
	//1235
	//45678
	return allarr
}

//快速排序
func QuickSortCallMerge(arr []int) []int {
	if len(arr) < 100 {
		return BinSearchSortMerge(arr)
	} else {
		QuickSortMerge(arr, 0, len(arr)-1)
		return arr
	}
}

//二分插入排序
func BinSearchSortMerge(mylist []int) []int {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindindexMid(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                          //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//快速排序，递归
func QuickSortMerge(arr []int, left int, right int) {
	if right-left < 100 {
		BinSearchSortIndexMerge(arr, left, right) //调用插入排序对于制定段排序
	} else {

		SwapMerge(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr[left]                                 //备份中间值
		It := left                                         // arr[left+1,lt]  <vdata  lt++
		gt := right + 1                                    //arr[gt...right]>vdata   gt--
		i := left + 1                                      // arr[lt+1, i] ==vdata   i++
		for i < gt {                                       //循环到重合
			if arr[i] < vdata {
				SwapMerge(arr, i, It+1) //移动小于的地方
				It++
				i++
			} else if arr[i] > vdata { //吧最右边大于4的数字与最左边小于4的数交换
				SwapMerge(arr, i, gt-1)
				gt--
			} else {
				i++ //相等
			}
		}
		SwapMerge(arr, left, It)
		QuickSortMerge(arr, left, It-1) //递归处理左边
		QuickSortMerge(arr, gt, right)  //递归处理右边
	}
}

//数据交换
func SwapMerge(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//对制定数据段排序
func BinSearchSortIndexMerge(mylist []int, start int, end int) []int {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMidMerge(mylist, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                   //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}
}

//二分查找
func FindindexMidMerge(list []int, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list[start] < list[cur] {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list[mid] > list[cur] {
		return FindindexMidMerge(list, start, mid, cur)
	} else {
		return FindindexMidMerge(list, mid+1, end, cur)
	}
}

//造数组
func makeArrMerge() []int {
	var length = 101
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)
	return list
}
