package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sync"

	"strconv"
)

// a b  a  x   a
//a   3       //a  2
//b   1       //b 1
//x   1     //z   3

// a 5
// b  2
// x  1
//z  3

func LoadFile() map[string]int {
	mymap := make(map[string]int)

	fi, err := os.Open("Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\all163_3.6pass.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return nil
	}
	defer fi.Close()
	br := bufio.NewReader(fi)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//计次数，存在+1，不存在次数为1
		if v, ok := mymap[string(line)]; ok {
			mymap[string(line)] = v + 1
		} else {
			mymap[string(line)] = 1
		}

	}

	return mymap
}

type Pass struct {
	PassWord string
	Times    int
}

func FindindexMid(list []Pass, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list[start].Times > list[cur].Times {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list[mid].Times < list[cur].Times {
		return FindindexMid(list, start, mid, cur)
	} else {
		return FindindexMid(list, mid+1, end, cur)
	}

}
func BinSearchSort(mylist []Pass) []Pass {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindindexMid(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                          //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}

//对制定数据段排序
func BinSearchSortIndex(mylist []Pass, start int, end int) []Pass {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMid(mylist, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                              //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}

}
func QuickSortCall(arr []Pass) []Pass {
	if len(arr) < 10 {
		return BinSearchSort(arr)
	} else {
		QuickSort(arr, 0, len(arr)-1)
		return arr
	}
}

//4 123 697
//123  4  697

//数据交换
func Swap(arr []Pass, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//快速排序，递归
func QuickSort(arr []Pass, left int, right int) {
	if right-left < 10 {
		BinSearchSortIndex(arr, left, right) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		Swap(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr[left]                            //备份中间值
		It := left                                    // arr[left+1,lt]  <vdata  lt++
		gt := right + 1                               //arr[gt...right]>vdata   gt--
		i := left + 1                                 // arr[lt+1, i] ==vdata   i++

		//	4 7 8 9  4 1 2  3
		//  i=1 vdata=4
		//	4 3 8 9  4 1 2 7
		//  i=1 vdata=4
		//	4 3 8 9  4 1 2 7
		//  i=1 vdata=4
		//	4 3 2 9  4 1 8 7
		//  i=2 vdata=4
		//	4 3 2 1  4   9 8 7
		//  i=2 vdata=4
		//1 3 2     4 4      9 8 7

		//	4 7 8 9  4  5 6
		//  //	44    7 8 9   5 6

		for i < gt { //循环到重合
			if arr[i].Times > vdata.Times {
				Swap(arr, i, It+1) //移动小于的地方
				It++
				i++

			} else if arr[i].Times < vdata.Times { //吧最右边大于4的数字与最左边小于4的数交换
				Swap(arr, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		Swap(arr, left, It)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			QuickSort(arr, left, It-1) //递归处理左边
			wg.Done()
		}()
		go func() {
			QuickSort(arr, gt, right) //递归处理右边
			wg.Done()
		}()
		wg.Wait()

	}
}

func main() {
	mymap := LoadFile()
	fmt.Println("读取完成", len(mymap)) //统计数量
	var N int = len(mymap)
	alldata := make([]Pass, N, N)
	i := 0
	//map迁移到了mymap
	for k, v := range mymap {
		alldata[i].PassWord = k
		alldata[i].Times = v
		i++
	}
	fmt.Println("迁移完成", len(alldata)) //统计数量
	//服务器优化
	mymap = nil
	runtime.GC()
	debug.FreeOSMemory()
	//排序
	alldata = QuickSortCall(alldata)
	fmt.Println("排序完成", len(alldata))
	//保存结果
	savefile, _ := os.Create("Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\all163_3.6passtimes.txt")
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //写入
	for i := 0; i < len(alldata); i++ {
		fmt.Fprintln(save, alldata[i].PassWord+" # "+strconv.Itoa(alldata[i].Times))
	}
	fmt.Println("保存完成", len(alldata))
	save.Flush()

}
