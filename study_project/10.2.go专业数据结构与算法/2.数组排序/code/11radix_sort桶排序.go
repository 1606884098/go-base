package main

import "fmt"

//银行，每人都有存款
//100 1000  10000
//99,    799 123  246,     1650  ,2370
//银行存款，金额固定，不固定

//800万高考考生  分数排序 0-750 1500
//10亿人，身高排序   0-300
// 687 3   688  6 689  8 690 -------------
//
func SelectSortMaxx(arr []int) int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0] //一个元素的数组，直接返回
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max { //任何一个比我的大的数，最大的
				max = arr[i]
			}
		}
		return max
	}

}

func RadixSort(arr []int) []int {
	max := SelectSortMaxx(arr) //寻找数组极大值 99991
	for bit := 1; max/bit > 0; bit *= 10 {
		//按照数量级分段
		arr = BitSort(arr, bit) //每次处理一个级别的排序
		fmt.Println(arr)
	}
	return arr
}
func BitSort(arr []int, bit int) []int {
	length := len(arr)           //数组长度
	bitcounts := make([]int, 10) //统计长度0,1,2,3,4,5,6,7,8,9
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 //分层处理，bit=1000的，三位数不参与排序了，bit=10000的四位数不参与排序
		bitcounts[num]++           //统计余数相等个数
	}
	fmt.Println(bitcounts)
	//  0 1 2 3  4 5
	//  1 0 3 0  0  1
	//  1 1 4 4  4  5
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1] //叠加，计算位置
	}
	fmt.Println(bitcounts)

	tmp := make([]int, 10) //开辟临时数组
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitcounts[num]-1] = arr[i] //计算排序的位置
		bitcounts[num]--
	}

	for i := 0; i < length; i++ {
		arr[i] = tmp[i] //保存数组
	}
	return arr

}

func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 111011}
	//11  222  7123
	//91
	//

	//fmt.Println(SelectSortMax(arr))
	fmt.Println(RadixSort(arr))

}
