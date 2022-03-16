package main

import "fmt"

//7亿女的，
// 身高排序
// 70-270
//70-270
//
// 1    2       3        -----
//70 ，71，72，73，--     165       -270

//0-500

//0-150

//幼儿园，小学，初中，中专，高中，大专，本科，硕士，博士
// 0，1，2，3，4，5，6，7

func main() {
	var arr [3][]int
	//012
	myarr := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 3}
	for i := 0; i < len(myarr); i++ {
		arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
	}
	fmt.Println(arr)
	//写入。

}
