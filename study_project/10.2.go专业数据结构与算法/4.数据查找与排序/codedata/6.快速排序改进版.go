package main

import (
	"fmt"
	"math/rand"
)

//12345
//1 left   5right

func SortForMerge(arr []int, left int, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] //备份数据
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- { //定位
			arr[j] = arr[j-1] //数据往后移动
		}
		arr[j] = temp //插入
	}
}

func swap(arr []int, i int, j int) { //数据交换
	arr[i], arr[j] = arr[j], arr[i]
}

//3       1 8 2 3  9 7
//2       1 3   3  897

//递归快速排序
func QuickSortX(arr []int, left int, right int) {
	if right-left < 2 { //数组剩下3个数，直接插入排序
		SortForMerge(arr, left, right)
	} else {
		//随机找一个数字，放在第一个位置
		swap(arr, left, rand.Int()%(right-left+1)+left)
		vdata := arr[left] //坐标数组，比我小，左边，比我大右边
		lt := left         // arr [left+1,  lt] <vata
		gt := right + 1    //arr[gt...  right] >vata
		i := left + 1      //arr[lt+1,...i] ==vdata
		for i < gt {
			if arr[i] < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++

			} else if arr[i] > vdata {
				swap(arr, i, gt-1) //移动到大于的地方
				gt--

			} else {
				i++
			}
		}
		swap(arr, left, lt)         //交换头部位置
		QuickSortX(arr, left, lt-1) //递归处理小于那一段
		QuickSortX(arr, gt, right)  //递归处理大于那一段

	}

}

//快速排序核心程序
func QuicksortPlus(arr []int) {
	QuickSortX(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println("未排序", arr)
	QuicksortPlus(arr)
	fmt.Println("已排序", arr)
}
