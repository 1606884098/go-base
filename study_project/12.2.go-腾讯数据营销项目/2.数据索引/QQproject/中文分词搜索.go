package main

import (
	"bufio"
	"fmt"
	"github.com/wangbin/jiebago"
	"io"
	"os"
	"strings"
)

var seg jiebago.Segmenter

func init() {
	seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
	for word := range ch {
		fmt.Printf(" %s /", word)
	}
	fmt.Println()
}

func main() {
	fmt.Print("【全模式】：")
	print(seg.CutAll("软件工程"))

	fmt.Print("【精确模式】：")
	print(seg.Cut("软件工程", false))

	fmt.Print("【新词识别】：")
	print(seg.Cut("软件工程", true))

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("软件工程", true))
}

//搜索字符串contains
//软件工程
//软件  工程
//软件
//工程
func main1() {
	fi, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\QQqunall.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close() //延迟关闭文件
	i := 0
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break //t跳出循环
		}
		linestr := string(line) //读取，转化为字符串
		print(seg.CutAll("我来到北京清华大学"))
		if strings.Contains(linestr, "软件工程") {
			fmt.Println(linestr)
		}
		//保存到文件，list

		if i%100000 == 0 {
			fmt.Println(i)
		}
		i++

	}

}
