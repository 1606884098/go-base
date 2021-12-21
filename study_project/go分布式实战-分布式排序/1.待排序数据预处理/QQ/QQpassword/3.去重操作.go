package main

import "fmt"
import "os"
import "bufio"
import (
	"io"
	"strconv"
)

func ReadfiletoArr(arr []string) {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpasswordsort.txt"
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
	const N = 156783175
	dataarr := make([]string, N+1, N+1)
	ReadfiletoArr(dataarr) //读取数据到数组

	//写入
	savefile, err := os.Create("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpasswordsorttimes.txt")
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
		for i+1 < length && dataarr[i] == dataarr[i+1] { //没有冒出数据的极限，
			i++
			times += 1
		}
		fmt.Fprintln(save, password+" # "+strconv.Itoa(times))

		//fmt.Println(password,times)

		i++
	}
	save.Flush()

}
