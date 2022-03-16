package main

import "fmt"

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		fmt.Println("head", arr)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
		fmt.Println("change", arr)
	}
	return arr
}

//1928
//9128
//128 9
//128 9
//12  89

func HeapSortMax(arr []int, length int) []int {
	//length:=len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //二叉树深度 2   ,遍历所有的中央节点，0，1，2
		for i := depth; i >= 0; i-- { //2,1,0
			topmax := i //嘉定最大的在i的位置
			leftchild := 2*i + 1
			rightchild := 2*i + 2                                      //左右节点
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
				topmax = leftchild
			}
			if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
				topmax = rightchild
			}
			if topmax != i {
				arr[i], arr[topmax] = arr[topmax], arr[i] //交换数据
			}

		}
		return arr
	}

}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5}
	fmt.Println(arr)
	fmt.Println(HeapSort(arr))
}
