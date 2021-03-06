package main

import "fmt"
import "os"
import "bufio"
import (
	"io"
	"strconv"
)

func main1() {
	mylist := []string{"1111", "222", "222", "222", "333", "333", "4444", "5555", "6666", "6666"}
	length := len(mylist)
	i := 0
	for i < length {
		times := 1
		password := mylist[i]
		for i+1 < length && mylist[i] == mylist[i+1] { //没有越界，
			i++
			times += 1 //统计出现次数
		}

		fmt.Println(password, times)
		i++
	}
}

func ReadfiletoArr11(arr []string) {
	path := "F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本排序后.txt"
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

func main() {
	const N = 6428632
	dataarr := make([]string, N+1, N+1)
	ReadfiletoArr11(dataarr) //读取数据到数组

	//写入
	savefile, err := os.Create("F:\\moni\\5.0go分布式实战-分布式排序\\1.待排序数据预处理\\csdn密码本排序后去重.txt")
	defer savefile.Close()
	if err != nil {
		fmt.Println("写入失败")
		return
	}
	save := bufio.NewWriter(savefile) //写入

	//去重写入
	length := len(dataarr)
	i := 0
	for i < length {
		times := 1
		password := dataarr[i]
		for i+1 < length && dataarr[i] == dataarr[i+1] { //没有越界，
			i++
			times += 1
		}
		fmt.Fprintln(save, password+" # "+strconv.Itoa(times)) //写入文件中

		//fmt.Println(password,times)

		i++
	}
	save.Flush()

}
