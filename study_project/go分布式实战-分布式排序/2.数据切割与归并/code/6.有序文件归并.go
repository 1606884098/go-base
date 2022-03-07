package main

import (
	"bufio"
	"io"
	"os"

	"fmt"
)

func GetLineNumberTestTwo(filepath string) int {
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

func MergreTestTwo(path1, path2, pathsave string) {
	length1 := GetLineNumberTestTwo(path1)
	length2 := GetLineNumberTestTwo(path2) //文件的长度
	fi1, _ := os.Open(path1)
	fi2, _ := os.Open(path2)
	defer fi1.Close()
	defer fi2.Close()
	br1 := bufio.NewReader(fi1)
	br2 := bufio.NewReader(fi2) //读取文件

	savefile, _ := os.Create(pathsave)
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //保存
	i, j := 0, 0
	line1, _, _ := br1.ReadLine()
	line2, _, _ := br2.ReadLine()

	for i < length1 && j < length2 {
		if string(line1) < string(line2) {
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++

		} else if string(line1) > string(line2) {
			fmt.Fprintln(save, string(line1))
			line1, _, _ = br1.ReadLine()
			i++

		} else {
			fmt.Fprintln(save, string(line1))
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++
			line1, _, _ = br1.ReadLine()
			i++
		}
	}
	for i < length1 {
		fmt.Fprintln(save, string(line1))
		line1, _, _ = br1.ReadLine()
		i++
	}
	for j < length2 {
		fmt.Fprintln(save, string(line2))
		line2, _, _ = br2.ReadLine()
		j++
	}
	save.Flush()
}

func main() {
	path1 := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\datasort\\163-0.txt"
	path2 := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\datasort\\163-1.txt"
	pathsave := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\datasort\\last.txt"
	MergreTestTwo(path1, path2, pathsave)
}
