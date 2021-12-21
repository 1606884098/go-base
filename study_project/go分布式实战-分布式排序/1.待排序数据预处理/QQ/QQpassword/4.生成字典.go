package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	//"github.com/hashicorp/vault/helper/password"
)

func GetLineNumber(filepath string) int {
	fi, err := os.Open(filepath)
	defer fi.Close() //打开文件
	if err != nil {
		return -1
	}
	i := 0 //统计行数
	br := bufio.NewReader(fi)
	for {
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break //文件结束，末尾
		}
		i++ //统计行数
		if i%5000000 == 0 {
			fmt.Println(i)
		}
	}
	return i

}
func main1() {
	const N = 117325850
	fmt.Println(GetLineNumber("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpasswordsorttimes.txt"))
}

type CSDNpassword struct {
	password string //密码
	times    int    //次数
}

func ReadfiletoArr(arr []CSDNpassword) {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpasswordsorttimes.txt"
	passfile, err := os.Open(path) //打开文件
	defer passfile.Close()
	if err != nil {
		fmt.Println("密码文件打开失败")
		return
	}
	br := bufio.NewReader(passfile) //读取文件对象
	i := 0
	for {
		line, _, end := br.ReadLine() //每次读取一行
		if end == io.EOF {
			break //文件结束跳出死循环
		}
		linestr := string(line)
		linelist := strings.Split(linestr, " # ")
		arr[i].password = linelist[0]               //读取密码，复制数组结构体元素之一
		arr[i].times, _ = strconv.Atoi(linelist[1]) //读取密码次数，复制数组结构体元素之一
		//arr[i]=string(line)//读取每一行
		i++

	}
}

func WritetoFile(arr []CSDNpassword) {
	savefile, err := os.Create("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpasswordlast.txt")
	defer savefile.Close()
	if err != nil {
		fmt.Println("写入失败")
		return
	}
	save := bufio.NewWriter(savefile) //写入
	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(save, arr[i].password+" # "+strconv.Itoa(arr[i].times))
	}
	save.Flush() //写入

}

func FindindexMid(list []CSDNpassword, start int, end int, cur int) int {
	//对比当前位置与需要排序的元素大小，返回较大值的位置
	if start >= end {
		if list[start].times > list[cur].times {
			return cur
		} else {
			return start
		}
	}
	mid := (start + end) / 2 //取得中间值

	//二分查找递归
	if list[mid].times < list[cur].times {
		return FindindexMid(list, start, mid, cur)
	} else {
		return FindindexMid(list, mid+1, end, cur)
	}

}
func BinSearchSort(mylist []CSDNpassword) []CSDNpassword {
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
func BinSearchSortIndex(mylist []CSDNpassword, start int, end int) []CSDNpassword {
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
func QuickSortCall(arr []CSDNpassword) []CSDNpassword {
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
func Swap(arr []CSDNpassword, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//快速排序，递归
func QuickSort(arr []CSDNpassword, left int, right int) {
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
			if arr[i].times > vdata.times {
				Swap(arr, i, It+1) //移动小于的地方
				It++
				i++

			} else if arr[i].times < vdata.times { //吧最右边大于4的数字与最左边小于4的数交换
				Swap(arr, i, gt-1)
				gt--

			} else {
				i++ //相等
			}
		}
		Swap(arr, left, It)

		QuickSort(arr, left, It-1) //递归处理左边
		QuickSort(arr, gt, right)  //递归处理右边

	}
}
func main() {
	//fmt.Println( GetLineNumber("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\passwordsorttimes.txt"))
	const N = 117325850
	CSDNarr := make([]CSDNpassword, N, N)

	//读取数据到  CSDNarr
	ReadfiletoArr(CSDNarr)
	fmt.Println("读取完成")
	//排序-从大到小
	CSDNarr = QuickSortCall(CSDNarr)
	fmt.Println("排序完成")

	//数据写入结果
	WritetoFile(CSDNarr)
	fmt.Println("写入完成")

	arrx := []int{1, 2, 3, 4} //分配一片内存
	fmt.Println(arrx)
	arrx = nil //内存标记不用
	runtime.GC()

}
