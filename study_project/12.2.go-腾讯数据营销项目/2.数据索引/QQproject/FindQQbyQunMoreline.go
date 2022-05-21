package main

import "fmt"
import "strconv"
import "os"
import (
	"strings"
)

//import "encoding/binary"

const N = 1449403409

func bin_searchQQ(findex, fii *os.File, QQqun int) []string {
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
		//fmt.Println("data",data)

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
		//fmt.Println("midstr",midstr)
		midlist := strings.Split(midstr, " # ")
		//fmt.Println("midlist",len(midlist),midlist[0],midlist[1],midlist[2])
		var midQQ int
		fmt.Sscanf(midlist[2], "%d", &midQQ) //数组载入内存
		//midQQ,_:=strconv.Atoi(midlist[2])
		//fmt.Println("midQQ",midQQ)

		if midQQ > QQqun {
			high = mid - 1
		} else if midQQ < QQqun {
			low = mid + 1
		} else {
			//mid//找到中间，
			mystrlist := make([]string, 0, 0)
			mystrlist = append(mystrlist, midstr)
			//向上循环
			tmp_up := mid
			for {
				tmp_up -= 1
				if tmp_up < low {
					break
				}
				findex.Seek(int64(tmp_up*15), 0)
				bx := make([]byte, 15, 15)
				_, _ = findex.Read(bx)
				var data int64
				fmt.Sscanf(string(bx), "%15d", &data)

				fii.Seek(int64(data), 0)
				b := make([]byte, 50, 50)
				length, _ := fii.Read(b)
				var i int
				for i = 0; i < length-1; i++ {
					if b[i] == '\n' && i >= 5+3+3+3+6 {
						break
					}
				}
				upstr := string(b[:i])
				uplist := strings.Split(upstr, " # ")
				var upQQ int
				fmt.Sscanf(uplist[2], "%d", &upQQ) //数组载入内存
				if upQQ == midQQ {
					mystrlist = append(mystrlist, upstr)
				} else {
					break
				}

			}

			//向下循环
			tmp_down := mid
			for {
				tmp_down += 1
				if tmp_up > high {
					break
				}
				findex.Seek(int64(tmp_down*15), 0)
				bx := make([]byte, 15, 15)
				_, _ = findex.Read(bx)
				var data int64
				fmt.Sscanf(string(bx), "%15d", &data)

				fii.Seek(int64(data), 0)
				b := make([]byte, 50, 50)
				length, _ := fii.Read(b)
				var i int
				for i = 0; i < length-1; i++ {
					if b[i] == '\n' && i >= 5+3+3+3+6 {
						break
					}
				}
				downstr := string(b[:i])
				downlist := strings.Split(downstr, " # ")
				var downQQ int
				fmt.Sscanf(downlist[2], "%d", &downQQ) //数组载入内存
				if downQQ == midQQ {
					mystrlist = append(mystrlist, downstr)
				} else {
					break
				}

			}

			return mystrlist //找到
		}
	}
	return []string{}

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
		mystrlist := bin_searchQQ(findex, fii, QQqun)
		if len(mystrlist) == 0 {
			fmt.Println("没有找到")
		} else {
			fmt.Println("找到", len(mystrlist), "个")
			for i := 0; i < len(mystrlist); i++ {
				fmt.Println(mystrlist[i])
			}
		}

	}
}
