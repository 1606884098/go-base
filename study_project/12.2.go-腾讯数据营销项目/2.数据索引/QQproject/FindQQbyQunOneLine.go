package main

import "fmt"
import "strconv"
import "os"
import "strings"

//import "encoding/binary"

const N = 1449403409

func bin_searchQQ(findex, fii *os.File, QQqun int) string {
	low := 0      //最下方
	high := N - 1 //最上方

	for low <= high { //循环的终止条件
		//fmt.Println(arr[low:high])
		mid := (low + high) / 2
		//fmt.Println("mid",mid)

		findex.Seek(int64(mid*15), 0)
		bx := make([]byte, 15, 15)
		_, _ = findex.Read(bx)
		var data int64
		fmt.Sscanf(string(bx), "%15d", &data) //数组载入内存

		//bx:=make([]byte,15,15)
		//num,err:=fmt.Sscanf(string(b),"%10d",&data)//数组载入内存
		//_,_=findex.Read(bx)
		//var data uint64
		//_,_=fmt.Sscanf(string(bx),"%15u",&data)
		//_,_=fmt.Sscanf(string(b),"%10d",&data)//数组载入内存
		fmt.Println("data", data)

		fii.Seek(int64(data), 0)
		b := make([]byte, 50, 50)
		length, _ := fii.Read(b)
		var i int
		for i = 0; i < length-1; i++ {
			if b[i] == '\n' && i >= 5+3+3+3+6 {
				break
			}
		}
		midstr := string(b[:i])
		fmt.Println("midstr", midstr)
		midlist := strings.Split(midstr, " # ")
		fmt.Println("midlist", len(midlist), midlist[0], midlist[1], midlist[2])
		var midQQ int
		fmt.Sscanf(midlist[2], "%d", &midQQ) //数组载入内存
		//midQQ,_:=strconv.Atoi(midlist[2])
		fmt.Println("midQQ", midQQ)

		if midQQ > QQqun {
			high = mid - 1
		} else if midQQ < QQqun {
			low = mid + 1
		} else {
			return midstr //找到
		}
	}
	return ""

}

func main() {
	findex, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\qun_data\\qun_qq_index.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer findex.Close() //延迟关闭文件
	fii, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\qun_data\\qun_name_qq.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fii.Close() //延迟关闭文件

	for {
		var input string
		fmt.Scanln(&input)
		QQqun, _ := strconv.Atoi(input)
		mystr := bin_searchQQ(findex, fii, QQqun)
		if mystr == "" {
			fmt.Println("没有找到")
		} else {
			fmt.Println("找到", mystr)
		}

	}
}

func main11() {
	fi, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\qun_data\\qun_qq_index.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	for i := 0; i < 10; i++ {
		b := make([]byte, 15, 15)
		_, _ = fi.Read(b)
		var data int
		num, err := fmt.Sscanf(string(b), "%15d", &data) //数组载入内存
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(num, data)
	}
}
