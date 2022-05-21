package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type QQ struct {
	QQnum    uint64
	PassWord string
}

func LoadfiletoMEM(N int) []QQ {
	var arrlist []QQ = make([]QQ, N, N)
	filepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"
	file, _ := os.Open(filepath) //打开文件
	br := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		linelist := strings.Split(linestr, "----")
		QQnum, _ := strconv.Atoi(linelist[0])
		arrlist[i].QQnum = uint64(QQnum)
		arrlist[i].PassWord = linelist[1]

		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println(i)
	return arrlist

}

func BinSearch(myarr []QQ, QQ uint64) int {
	low := 0
	high := len(myarr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(myarr[mid].QQnum, QQ)
		if myarr[mid].QQnum > QQ {
			high = mid - 1

		} else if myarr[mid].QQnum < QQ {
			low = mid + 1

		} else {
			return mid
		}
	}
	return -1

}

func main() {
	fmt.Println("读取行数")
	N := 156774123
	fmt.Println("载入数据到内存")
	myarr := LoadfiletoMEM(N)

	for {
		var QQ uint64
		fmt.Scanf("%d", &QQ)
		id := BinSearch(myarr, QQ)
		if id == -1 {
			fmt.Println("没有找到")
		} else {
			fmt.Println("找到", myarr[id])
		}
	}

}
