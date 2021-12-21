package main

import "fmt"

// 1         9   2  8  6    3  5
//1 9      2  8  6  3  5
//1  2 9      8  6 3  5
///1  2  8  9       6 3  5
///1  2 6   8  9         3  5

///1  2 3  6   8  9           5

///1  2 3  5  6   8  9

func SimpleInsert(arr []int) []int {
	backup := arr[3]
	j := 3 - 1                      //上一个位置开始循环
	for j >= 0 && backup < arr[j] { //从前往后移动
		arr[j+1] = arr[j] //从前往后移动
		//1,19,29,8,3,7,4,6,5,10
		//1,19,19,29,3,7,4,6,5,10
		//1,8,19,29,3,7,4,6,5,10
		j--
		fmt.Println(arr)
	}
	arr[j+1] = backup
	return arr
}

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 1; i < len(arr); i++ {
			backup := arr[i]
			j := i - 1                      //上一个位置开始循环
			for j >= 0 && backup < arr[j] { //从前往后移动
				arr[j+1] = arr[j] //从前往后移动
				//1,19,29,8,3,7,4,6,5,10
				//1,19,19,29,3,7,4,6,5,10
				//1,8,19,29,3,7,4,6,5,10
				j--
				fmt.Println(arr)
			}
			arr[j+1] = backup
		}

		return arr
	}

}

func main() {
	arr := []int{10, 19, 29, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SimpleInsert(arr))
	fmt.Println(InsertSort(arr))
}
