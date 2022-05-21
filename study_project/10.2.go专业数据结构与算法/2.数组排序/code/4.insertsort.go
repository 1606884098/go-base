package main

import "fmt"

////       1        9 2 8  7 6  4  5
////       1 9          2 8  7 6  4  5
////       1 2  9           8  7 6  4  5
////       1 2  8  9            7 6  4  5
////       1 2   7  8  9            6  4  5
////       1 2  6 7  8  9              4  5

//1      9         2 8  7 6  4  5
//1 9  9
//1 2  9
//1  19 29    8
// 1  19 29  29
// 1  19 19  29
// 1  8 19  29
func InsertTest(arr []int) []int {
	backup := arr[3]
	j := 3 - 1 //上一个位置循环找到位置插入
	for j >= 0 && backup < arr[j] {
		arr[j+1] = arr[j] //从前往后移动
		j--
	}
	arr[j+1] = backup //插入
	return arr

}

func InsertSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 1; i < length; i++ { //跳过第一个
			backup := arr[i] //备份插入的数据
			j := i - 1       //上一个位置循环找到位置插入
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j] //从前往后移动
				j--
			}
			arr[j+1] = backup //插入
			fmt.Println(arr)
		}

		return arr

	}
}

func main() {
	arr := []int{1, 19, 29, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(InsertTest(arr))
	fmt.Println(InsertSort(arr))

}
