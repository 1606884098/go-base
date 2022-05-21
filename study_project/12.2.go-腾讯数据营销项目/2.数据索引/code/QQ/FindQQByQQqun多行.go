package main

import (
	"fmt"
	"os"
	"strings"
	//"strconv"
	//"encoding/binary"
)

func BinSearchdisk(file *os.File, fileindex *os.File, QQqun int) []string {
	const N = 156774124
	low := 0
	high := N - 1

	for low < high {
		mid := (low + high) / 2

		fileindex.Seek(int64(mid*15), 0)
		bx := make([]byte, 15, 15)
		_, _ = fileindex.Read(bx)
		var depdata int64
		fmt.Sscanf(string(bx), "%15d", &depdata)

		file.Seek(int64(depdata), 0) //跳到指定行
		b := make([]byte, 150, 150)
		length, _ := file.Read(b) //读取
		var endpos int
		for i := 0; i < length-1; i++ {
			if b[i] == '\n' && i >= 5+3+3+3+6 {
				endpos = i
				break
			}
		}

		midstr := string(b[:endpos])
		fmt.Println("mid数据", midstr)
		midlist := strings.Split(midstr, " # ")
		//midQQ,_:=strconv.Atoi(midlist[2])
		var midQQ int
		fmt.Sscanf(midlist[2], "%d", &midQQ)

		if midQQ > QQqun {
			high = mid - 1
		} else if midQQ < QQqun {
			low = mid + 1
		} else {
			mylist := []string{}
			mylist = append(mylist, midstr) //第一个

			tmp_up := mid
			for {
				tmp_up -= 1
				if tmp_up < low {
					break
				}

				fileindex.Seek(int64(tmp_up*15), 0)
				bx := make([]byte, 15, 15)
				_, _ = fileindex.Read(bx)
				var depdata int64
				fmt.Sscanf(string(bx), "%15d", &depdata)

				file.Seek(int64(depdata), 0) //跳到指定行
				b := make([]byte, 150, 150)
				length, _ := file.Read(b) //读取
				var endpos int
				for i := 0; i < length-1; i++ {
					if b[i] == '\n' && i >= 5+3+3+3+6 {
						endpos = i
						break
					}
				}
				upstr := string(b[:endpos]) //获取字符串
				uplist := strings.Split(upstr, " # ")
				var upQQ int
				fmt.Sscanf(uplist[2], "%d", &upQQ)
				if upQQ == midQQ {
					mylist = append(mylist, upstr)
				} else {
					break
				}

			}

			tmp_down := mid
			for {
				tmp_down += 1
				if tmp_down > high {
					break
				}

				fileindex.Seek(int64(tmp_down*15), 0)
				bx := make([]byte, 15, 15)
				_, _ = fileindex.Read(bx)
				var depdata int64
				fmt.Sscanf(string(bx), "%15d", &depdata)

				file.Seek(int64(depdata), 0) //跳到指定行
				b := make([]byte, 150, 150)
				length, _ := file.Read(b) //读取
				var endpos int
				for i := 0; i < length-1; i++ {
					if b[i] == '\n' && i >= 5+3+3+3+6 {
						endpos = i
						break
					}
				}
				downstr := string(b[:endpos]) //获取字符串
				downlist := strings.Split(downstr, " # ")
				var upQQ int
				fmt.Sscanf(downlist[2], "%d", &upQQ)
				if upQQ == midQQ {
					mylist = append(mylist, downstr)
				} else {
					break
				}
			}

			return mylist
		}

	}
	return []string{}
}

func main() {
	path := "Z:\\J\\洗币\\社会工程学\\newQQ\\qun_data\\qun_name_qq.txt"
	file, _ := os.Open(path) //打开文件
	file.Seek(0, 0)          //移动到文件开头

	pathindex := "Z:\\J\\洗币\\社会工程学\\newQQ\\qun_data\\qun_qq_index.txt"
	fileindex, _ := os.Open(pathindex) //打开文件
	fileindex.Seek(0, 0)               //移动到文件开头

	for {
		var QQqun int
		fmt.Scanf("%d", &QQqun)
		mylist := BinSearchdisk(file, fileindex, QQqun)
		if len(mylist) == 0 {
			fmt.Println("没有找到")
		} else {
			for i := 0; i < len(mylist); i++ {
				fmt.Println("找到", mylist[i])
			}
		}

	}
	file.Close()
	fileindex.Close()
}
