package main

import "fmt"

//4 8 1 2 3  7 9
//4
//123  4   879

//123  4   //789

//数组<100 ,二分插入排序大于快速排序
//if 数组长度<100
//二分插入排序
//>=100快速排序

func QuickSort1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		splitdata := arr[0]          //第一个数据
		low := make([]int, 0, 0)     //比我小
		high := make([]int, 0, 0)    //比我大
		mid := make([]int, 0, 0)     //与我一样大
		mid = append(mid, splitdata) //加入一个

		for i := 1; i < len(arr); i++ {
			if arr[i] > splitdata {
				low = append(low, arr[i])
			} else if arr[i] < splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort1(low), QuickSort1(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func main() {
	arr := []int{12324, 19, 111, 237, 6, 5, 10, 11, 223}
	fmt.Println(QuickSort1(arr))

}
