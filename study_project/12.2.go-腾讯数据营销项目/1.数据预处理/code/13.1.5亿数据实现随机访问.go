package main

import "os"
import (
	"encoding/binary"
	"fmt"
)

func main() {
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"
	file, _ := os.Open(path) //打开文件
	file.Seek(0, 0)          //移动到文件开头

	pathindex := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodindex.txt"
	fileindex, _ := os.Open(pathindex) //打开文件
	fileindex.Seek(0, 0)               //移动到文件开头

	//lengthlist:=makeIndex()
	for {
		var linenum int
		fmt.Scanf("%d", &linenum)

		fileindex.Seek(int64(linenum*4), 0)
		bx := make([]byte, 4, 4)
		_, _ = fileindex.Read(bx)
		var depdata uint32 = uint32(binary.BigEndian.Uint32(bx)) //二进制专户为整数
		fmt.Println("linenum", linenum, "depdata", depdata)

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
		fmt.Println("显示数据", string(b[:endpos]))
	}
	fileindex.Close()
	file.Close()
}
