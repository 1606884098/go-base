package main

import "math/rand"
import "time"
import "fmt"

//1 9 2 8 3         7 4 6 5
//1 9 2       8 3         7 4      6 5
//1     9 2       8    3         7    4      6   5
//1     9     2       8    3         7    4      6   5
//19   2   38      47    56
//129    3478  56
//129  345678
//123456789

func main() {
	myarr := makeArrMergetoSort()
	fmt.Println(myarr)
	fmt.Println(MergeSortGrop(myarr))
}

//归并排序的简单归并
func MergeSimple(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
			result = append(result, right[j])
			j++
		}
	}
	for i < l {
		result = append(result, left[i])
		i++
	}
	for j < r {
		result = append(result, right[j])
		j++
	}
	return result
}

//递归调用归并
func MergeSortGrop(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	i := len(arr) / 2
	left := MergeSortGrop(arr[0:i]) //浪费空间
	right := MergeSortGrop(arr[i:])
	result := MergeSimple(left, right)
	return result
}

func makeArrMergetoSort() []int {
	var length = 30
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	//fmt.Println(list)
	return list
}
