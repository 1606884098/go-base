package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	//"strings"
	//"github.com/hashicorp/vault/helper/password"
)

//将字符串数据，快速排序
func main() {
	//fmt.Println(GetLineNumber("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\password.txt"))
	const N = 6428632
	dataarr := make([]string, N+1, N+1)
	//读取文件数据到数组
	ReadfiletoArr(dataarr)
	fmt.Println("读取完成，成功将数据加载到内存中")
	//排序
	dataarr = QuickSortCallString(dataarr)
	fmt.Println("排序完成")
	//写入到新的文件
	WritetoFile(dataarr)
	fmt.Println("将排序后的数据写入文件完成！")
}

//获取文件的行数
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
	}
	return i
}

//将文件的密码读取到内存数组中
func ReadfiletoArr(arr []string) {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\password.txt"
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
		arr[i] = string(line) //读取每一行
		i++
	}
}

func QuickSortCallString(arr []string) []string {
	if len(arr) < 100 {
		return BinSearchSortString(arr)
	} else {
		QuickSort(arr, 0, len(arr)-1)
		return arr
	}
}

//二分插入排序发
func BinSearchSortString(mylist []string) []string {
	if len(mylist) <= 1 {
		return mylist
	} else {
		for i := 1; i < len(mylist); i++ {
			p := FindindexMidString(mylist, 0, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}
}

//二分查找
func FindindexMidString(list []string, start int, end int, cur int) int {
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
		return FindindexMidString(list, start, mid, cur)
	} else {
		return FindindexMidString(list, mid+1, end, cur)
	}
}

//快速排序，递归
func QuickSort(arr []string, left int, right int) {
	if right-left < 10 {
		BinSearchSortIndexString(arr, left, right) //调用插入排序对于制定段排序
	} else {
		Swap(arr, left, rand.Int()%(right-left)+left) //任何一个位置，交换到第一个
		vdata := arr[left]                            //备份中间值
		It := left                                    // arr[left+1,lt]  <vdata  lt++
		gt := right + 1                               //arr[gt...right]>vdata   gt--
		i := left + 1                                 // arr[lt+1, i] ==vdata   i++
		for i < gt {                                  //循环到重合
			if arr[i] < vdata {
				Swap(arr, i, It+1) //移动小于的地方
				It++
				i++

			} else if arr[i] > vdata { //吧最右边大于4的数字与最左边小于4的数交换
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

//对制定数据段排序
func BinSearchSortIndexString(mylist []string, start int, end int) []string {
	if end-start <= 1 {
		return mylist
	} else {
		for i := start + 1; i <= end; i++ {
			p := FindindexMidString(mylist, start, i-1, i) //0,0,  0,1,  0,2,   0,3
			if p != i {                                    //不等，插入
				for j := i; j > p; j-- {
					mylist[j], mylist[j-1] = mylist[j-1], mylist[j] //数据移动
				}
			}
		}
		return mylist
	}
}

//数据交换
func Swap(arr []string, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//将数据写入新文件中
func WritetoFile(arr []string) {
	savefile, err := os.Create("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\passwordsortQ.txt")
	defer savefile.Close()
	if err != nil {
		fmt.Println("写入失败")
		return
	}
	save := bufio.NewWriter(savefile) //写入
	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(save, arr[i])
	}
	save.Flush() //写入
}
