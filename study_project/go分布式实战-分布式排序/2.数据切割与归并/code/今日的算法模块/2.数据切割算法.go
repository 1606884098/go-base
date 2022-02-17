package main

import "fmt"

//100  10
//97  10     10  10  10.。。。 7

func evgSplit(num, N int) []int {
	arr := []int{}
	if num%N == 0 {
		for i := 0; i < N; i++ {
			arr = append(arr, num/N)
		}
	} else {
		evg := (num - num%N) / (N - 1) //97 -7 //9=10
		for i := 0; i < N-1; i++ {
			arr = append(arr, evg) //追加
			num -= evg
		}
		arr = append(arr, num) //7

	}
	return arr
}
func main() {
	fmt.Println(evgSplit(100, 10))
	fmt.Println(evgSplit(97, 10))
	fmt.Println(evgSplit(106, 10))
	fmt.Println(evgSplit(176, 10))
	fmt.Println(evgSplit(576, 10))
}
