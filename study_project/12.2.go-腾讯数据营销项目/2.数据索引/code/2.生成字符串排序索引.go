package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
)

const N = 156774123

func LoadfiletoMEM() []string {
	//const  N=1 5678 3175

	var arrlist []string = make([]string, N, N)
	filepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGood.txt"
	file, _ := os.Open(filepath) //打开文件
	br := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		arrlist[i] = string(line)
		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println(i)
	return arrlist

}
func makeindexArr(arrlist []string) []int {
	lengthlist := make([]int, N+1, N+1)
	lengthlist[0] = 0
	for i := 0; i < len(lengthlist)-1; i++ {
		lengthlist[i+1] = len(arrlist[i]) + 1 //1换行符，
	}
	for j := 0; j < len(lengthlist)-1; j++ {
		lengthlist[j+1] += lengthlist[j] //叠加

	}
	fmt.Println("索引生成了")
	return lengthlist

}

func FindindexMid(QQ []string, list2 []int, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		mystr1 := QQ[start]
		mylist1 := strings.Split(mystr1, "----")

		mystr2 := QQ[cur]
		mylist2 := strings.Split(mystr2, "----")

		if mylist1[0] > mylist2[0] {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	mystr1x := QQ[mid]
	mylist1x := strings.Split(mystr1x, "----")

	mystr2x := QQ[cur]
	mylist2x := strings.Split(mystr2x, "----")

	if mylist1x[0] < mylist2x[0] {
		return FindindexMid(QQ, list2, start, mid, cur)
	} else {
		return FindindexMid(QQ, list2, mid+1, end, cur)
	}

}
func BinSearchSort(QQ []string, mylist2 []int) []int {
	if len(mylist2) <= 1 {
		return mylist2
	} else {
		for i := 1; i < len(mylist2); i++ {
			p := FindindexMid(QQ, mylist2, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                               //不等，插入
				for j := i; j > p; j-- {
					mylist2[j], mylist2[j-1] = mylist2[j-1], mylist2[j] //数据移动
				}
			}
		}
		return mylist2
	}

}

//对制定数据段排序
func BinSearchSortIndex(QQ []string, mylist2 []int, start int, end int) []int {
	if end-start <= 1 {
		return mylist2
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMid(QQ, mylist2, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                   //不等，插入
				for j := i; j > p; j-- {
					mylist2[j], mylist2[j-1] = mylist2[j-1], mylist2[j] //数据移动
				}
			}
		}
		return mylist2
	}

}
func QuickSortCall(QQ []string, arr2 []int) []int {
	if len(arr2) <= 1 {
		return BinSearchSort(QQ, arr2)
	} else {
		QuickSort(QQ, arr2, 0, len(QQ)-1)
		return arr2
	}
}

//4 123 697
//123  4  697

//数据交换
func Swap(QQ []string, arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	QQ[i], QQ[j] = QQ[j], QQ[i]
}

//快速排序，递归
func QuickSort(QQ []string, arr2 []int, left int, right int) {
	if right-left < 1 {
		BinSearchSortIndex(QQ, arr2, left, right) //调用插入排序对于制定段排序
	} else {
		//快速排序写法
		//第一个，最后一个，随机抓取
		// 4 123  4 789  10
		//
		Swap(QQ, arr2, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := left                                      //备份中间值
		It := left                                         // arr[left+1,lt]  <vdata  lt++
		gt := right + 1                                    //arr[gt...right]>vdata   gt--
		i := left + 1                                      // arr[lt+1, i] ==vdata   i++

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
			mystr1x := QQ[i]
			mylist1x := strings.Split(mystr1x, "----")

			mystr2x := QQ[vdata]
			mylist2x := strings.Split(mystr2x, "----")

			if mylist1x[0] > mylist2x[0] {
				Swap(QQ, arr2, i, It+1) //移动小于的地方
				It++
				i++

			} else if mylist1x[0] < mylist2x[0] { //吧最右边大于4的数字与最左边小于4的数交换
				Swap(QQ, arr2, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		Swap(QQ, arr2, left, It)

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			QuickSort(QQ, arr2, left, It-1) //递归处理左边
			wg.Done()
		}()
		go func() {
			QuickSort(QQ, arr2, gt, right) //递归处理右边
			wg.Done()
		}()
		wg.Wait()

	}
}
func main() {
	arrlist := LoadfiletoMEM()
	fmt.Println("数据载入内存")
	lengthlist := makeindexArr(arrlist)
	fmt.Println("索引生成")
	lengthlist = QuickSortCall(arrlist, lengthlist) //排序
	//保存写入到文件
	fmt.Println("索引排序")

	savepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodindexSort.txt"
	savefile, _ := os.Create(savepath)
	wr := bufio.NewWriter(savefile)
	defer savefile.Close()
	for i := 0; i < len(lengthlist); i++ {
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(lengthlist[i]))
		//savefile.Write(b)
		wr.Write(b)
		if i%10000 == 0 {
			fmt.Println(i, lengthlist[i])
		}
	}
	wr.Flush()
	fmt.Println("索引保存完成")

}
