package main

import "fmt"

func ComoSort(arr []int) []int {
	length := len(arr)
	gap := length
	for gap > 1 {
		gap = gap * 10 / 13
		for i := 0; i+gap < length; i++ { //收缩，
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}

	}
	return arr
}

func main() {
	var array []int = []int{16, 8, 1, 24, 30}
	fmt.Println(array)
	ComoSort(array)
	fmt.Println(array)
}
