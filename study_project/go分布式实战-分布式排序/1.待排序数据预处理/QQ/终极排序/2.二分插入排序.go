package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makearr1() []int {
	var length = 10
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)
	return list
}

//list数组，start开始位置，end结束位置，cur当前
//  0  1  2   3  4  5 6 7 8
// 0  1   2  4   5 6 7 8 9           3   ->2,4
// start=0    mid=4    end=8
//      arr[mid]=5

// start=0    mid=2    end=4
//      arr[mid]=2
//start=3,   mid=3  end=4
//      arr[mid]=4
// return 3

func FindindexMid1(list []int, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list[start] < list[cur] {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list[mid] > list[cur] {
		return FindindexMid(list, start, mid, cur)
	} else {
		return FindindexMid(list, mid+1, end, cur)
	}

}

// cur=7 mid=  5 end= 9
//      start
// 1，2，4，5，6，7，8  ,9//  3

func main() {
	mylist := makearr1()
	for i := 1; i < len(mylist); i++ {
		//循环插入，寻找合适位置，
		p := FindindexMid1(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
		//1234
		//1230
		if p != i { //不等，插入
			//1230
			//0123
			//	temp:=mylist[i]//备份一下
			for j := i; j > p; j-- {
				mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
			}
			//mylist[p]=temp//填充
		} else {
			//123  4
		}
	}
	fmt.Println(mylist)
}

//123 4
