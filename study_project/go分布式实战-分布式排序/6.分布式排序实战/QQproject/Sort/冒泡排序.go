package main

import "fmt"

// 1 9  2 8  3 7  4  6  5
// 1 2  8 3  7  4 6 5  9

// 1 2 3 7   4  6 5  89

func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1] //冒泡法
		}
		fmt.Println(arr)
	}
	return arr[len(arr)-1]

}
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i] //冒泡法
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5}
	fmt.Println(arr)
	fmt.Println(BubbleSort(arr))
}

//[ 11  1 2 3 4 5 6 7 8 9]
