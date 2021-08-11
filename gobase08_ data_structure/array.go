package main

import "fmt"

func main() {
	/*var 数组变量名 [元素数量]Type*/
	var intArray [3]int = [3]int{1, 2, 3}
	//a:=[3]int{1,2,3}
	//d:=[5]int{2:10,4:9}指定下标来赋值  下标：值
	//b:=[...]int{1,2,3}//通过赋值来确定个数

	for i := 0; i < len(intArray); i++ { //普通遍历
		//从打印出来的地址可以看出是连续的地址空间，空间的大小跟类型有关
		fmt.Printf("index=%d,val=%d,address=%x\n", i, intArray[i], &intArray[i])
		/*		index=0,val=1,address=c0000a0120
				index=1,val=2,address=c0000a0128
				index=2,val=3,address=c0000a0130*/
	}

	for i, v := range intArray {
		fmt.Printf("index=%d,val=%d\n", i, v)
	}
}
