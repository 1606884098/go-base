package main

import "fmt"

//11,9,2,8,3,7,4,6,5,10
//9 11 2 8 3 7 4 6 5 10
//9 2 11 8 3 7 4 6 5 10
//9 2 8 11 3 7 4 6 5 10
//9 2 8 3 11 7 4 6 5 10
//9 2 8 3 7 11  4 6 5 10
//9 2 8 3 7 4  11 6 5 10
//9 2 8 3 7 4  6 11  5 10
//9 2 8 3 7 4  6  5 11 10
//9 2 8 3 7 4  6  5 10 11
//9 2 8 3 7 4  6  5 10

func BubbleFindMax(arr []int) int {
	length := len(arr) //求数组长度
	if length <= 1 {
		return arr[0]
	} else {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] { //两两比较
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		return arr[length-1]
	}
}

func Bubblesort(arr []int) []int {
	length := len(arr) //求数组长度
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个，不需要冒泡了
			isneedexchange := false
			for j := 0; j < length-1-i; j++ {
				if arr[j] > arr[j+1] { //两两比较
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isneedexchange = true
				}
			}
			if !isneedexchange {
				break
			}
			fmt.Println(arr)

		}

		return arr
	}
}

func main() {
	arr := []int{11, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(BubbleFindMax(arr))
	fmt.Println(Bubblesort(arr))

}
