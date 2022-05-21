package main

import "fmt"

// 8 1 4 2 9 5 3
// 1 4 2 8 5  39 //9沉底  正向冒泡一次
// 1 4 2 8 5  39
//  1 2 4 3 8  5 9 //2飘起来  反向冒泡，得到极大，极小
//

func cocktail(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ { //每次循环，正向冒泡一次，反向冒泡一次
		left := 0
		right := len(arr) - 1 //左边，右边

		for left <= right { //结束的条件
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--

		}
		fmt.Println(i, arr)
	}

	return arr

}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(cocktail(arr))

}
