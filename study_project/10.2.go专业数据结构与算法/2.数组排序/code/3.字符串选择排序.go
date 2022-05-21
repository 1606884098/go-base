package main

import (
	"fmt"
	"strings"
)

func SelectSortMaxString(arr []string) string {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0] //一个元素的数组，直接返回
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			if strings.Compare(arr[i], max) > 0 {
				max = arr[i]
			}
			/*
				if arr[i]>max{ //任何一个比我的大的数，最大的
					max=arr[i]
				}*/
		}
		return max
	}

}

func SelectSortString(arr []string) []string {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要挑选，
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				/*if arr[min]>arr[j]{
					min=j //保存极小值的索引
				}*/
				if strings.Compare(arr[min], arr[j]) < 0 {
					min = j
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
	arr := []string{"1c", "1a", "1b", "1x", "1z", "1m", "1n", "1d", "1f"}
	fmt.Println(SelectSortMaxString(arr))
	fmt.Println(SelectSortString(arr))

}
