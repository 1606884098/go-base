package main

import "os"
import "io"
import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//4037920
type CSDNpassword struct {
	password string //密码
	times    int    //次数
}

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

func ReadfiletoArr(arr []CSDNpassword) {
	path := "F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本排序后去重.txt"
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
func QuickSort(arr []CSDNpassword) []CSDNpassword {
	if len(arr) <= 1 {
		return arr
	} else {
		splitdata := arr[0]                //第一个数据
		low := make([]CSDNpassword, 0, 0)  //比我小
		high := make([]CSDNpassword, 0, 0) //比我大
		mid := make([]CSDNpassword, 0, 0)  //与我一样大
		mid = append(mid, splitdata)       //加入一个

		for i := 1; i < len(arr); i++ {
			if arr[i].times > splitdata.times {
				low = append(low, arr[i])
			} else if arr[i].times < splitdata.times {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}
func WritetoFile(arr []CSDNpassword) {
	savefile, err := os.Create("F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本排序后去重后排序.txt")
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
func main() {
	//fmt.Println( GetLineNumber("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\passwordsorttimes.txt"))
	const N = 4037920
	CSDNarr := make([]CSDNpassword, N, N)

	//读取数据到  CSDNarr
	ReadfiletoArr(CSDNarr)
	fmt.Println("读取完成")
	//排序-从大到小
	CSDNarr = QuickSort(CSDNarr)
	fmt.Println("排序完成")
	//数据写入结果
	WritetoFile(CSDNarr)
	fmt.Println("写入完成")

}
