package main

import "os"

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func BinSearchdisk(file *os.File, fileindex *os.File, QQ int) string {
	const N = 156774124
	low := 0
	high := N - 1

	for low < high {
		mid := (low + high) / 2

		fileindex.Seek(int64(mid*4), 0)
		bx := make([]byte, 4, 4)
		_, _ = fileindex.Read(bx)
		var depdata uint32 = uint32(binary.BigEndian.Uint32(bx)) //二进制专户为整数

		file.Seek(int64(depdata), 0) //跳到指定行
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

	pathindex := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodindex.txt"
	fileindex, _ := os.Open(pathindex) //打开文件
	fileindex.Seek(0, 0)               //移动到文件开头

	for {
		var QQ int
		fmt.Scanf("%d", &QQ)
		fmt.Println("显示数据", BinSearchdisk(file, fileindex, QQ))
	}

	file.Close()
	fileindex.Close()
}
