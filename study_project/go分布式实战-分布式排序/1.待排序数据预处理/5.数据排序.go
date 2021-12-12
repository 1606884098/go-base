package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strings"
	//"github.com/hashicorp/vault/helper/password"
)

func GetLineNumber1(filepath string) int {
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

func ReadfiletoArr1(arr []string) {
	path := "F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本.txt"
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

func WritetoFile1(arr []string) {
	savefile, err := os.Create("F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本排序后.txt")
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
func QuickSort11(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	} else {
		splitdata := arr[0]          //第一个数据
		low := make([]string, 0, 0)  //比我小
		high := make([]string, 0, 0) //比我大
		mid := make([]string, 0, 0)  //与我一样大
		mid = append(mid, splitdata) //加入一个

		for i := 1; i < len(arr); i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort11(low), QuickSort11(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func main() {
	const N = 6428632
	dataarr := make([]string, N+1, N+1)
	//读取文件数据到数组
	ReadfiletoArr1(dataarr)
	fmt.Println("读取完成1")
	//排序
	dataarr = QuickSort11(dataarr)
	fmt.Println("排序完成1")
	//写入到新的文件
	WritetoFile1(dataarr)
	fmt.Println("写入完成1")

	//fmt.Println(GetLineNumber1("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\password.txt"))
}
