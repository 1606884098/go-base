package main

import "fmt"

//topk
//百度，淘宝搜索，---日志
//关键字搜索次数最多 topk1000

//选择排序
//1，9，2，8，3，7，4，6，5
//9
//9  8
//9  8  7

//返回最大值
func SelectMax(arr []int) int {
	length := len(arr) //取得数组长度
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0] //嘉定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i] //始终存储最大
			}
		}
		return max

	}
}
func SelectSort(arr []int) []int {
	length := len(arr) //取得数组长度
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j //存储最小的
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
			fmt.Println(arr)
		}
		return arr

	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5} //数组排序
	fmt.Println(SelectSort(arr))
	//10  0000  0000

}
