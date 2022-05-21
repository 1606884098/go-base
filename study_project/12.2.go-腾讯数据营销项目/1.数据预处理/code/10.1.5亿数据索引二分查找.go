package main

import "os"
import "io"
import "bufio"
import (
	"fmt"
	"strconv"
	"strings"
)

func makeIndex() []int {
	N := 156774123
	indexdata := make([]int, N, N)
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"

	file, _ := os.Open(path) //打开文件
	br := bufio.NewReader(file)
	indexdata[0] = 0
	i := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		if i < N {
			indexdata[i] = len(linestr) + 1
		}
		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println("第一步读取长度索引", i)

	for j := 0; j < len(indexdata)-1; j++ {
		indexdata[j+1] += indexdata[j]
	}
	fmt.Println("第二部数组索引叠加完成")

	return indexdata

}

const N = 156774124

func BinSearch(file *os.File, lengthlist []int, QQ int) string {
	low := 0
	high := len(lengthlist) - 1

	for low < high {
		mid := (low + high) / 2

		file.Seek(int64(lengthlist[mid]), 0) //跳到指定行
		b := make([]byte, 12+4+16, 12+4+16)
		length, _ := file.Read(b) //读取
		var endpos int
		for i := 0; i < length-1; i++ {
			if b[i] == '\n' && i >= 5+4+6 {
				endpos = i
				break
			}
		}

		midstr := string(b[:endpos])
		fmt.Println("mid数据", midstr)
		midlist := strings.Split(midstr, "----")
		midQQ, _ := strconv.Atoi(midlist[0])

		if midQQ > QQ {
			high = mid - 1
		} else if midQQ < QQ {
			low = mid + 1
		} else {
			return midstr
		}

	}
	return "没有找到"
}

func main() {
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"
	file, _ := os.Open(path) //打开文件
	file.Seek(0, 0)          //移动到文件开头
	lengthlist := makeIndex()
	for {
		var QQ int
		fmt.Scanf("%d", &QQ)
		fmt.Println("显示数据", BinSearch(file, lengthlist, QQ))
	}

	file.Close()
}

func main1x() {
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"
	file, _ := os.Open(path) //打开文件
	file.Seek(0, 0)          //移动到文件开头
	lengthlist := makeIndex()
	for {
		var linenum int
		fmt.Scanf("%d", &linenum)
		file.Seek(int64(lengthlist[linenum]), 0) //跳到指定行
		b := make([]byte, 12+4+16, 12+4+16)
		length, _ := file.Read(b) //读取
		var endpos int
		for i := 0; i < length-1; i++ {
			if b[i] == '\n' && i >= 5+4+6 {
				endpos = i
				break
			}
		}
		fmt.Println("显示数据", string(b[:endpos]))
	}

	file.Close()
}
