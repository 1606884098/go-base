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

//N=86907937 行
func init() {
	seg.LoadDictionary("dict.txt")
}

func mainSearch(searchstr string) []string {

	filesavepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day7\\tmp\\"
	filelastpath := filesavepath + searchstr + ".txt"
	_, err := os.Stat(filelastpath)
	resultstring := []string{}
	if err == nil {
		//存在
		fmt.Println("文件存在")
		//resultstring=append(resultstring,)
		file, _ := os.Open(filelastpath) //打开文件
		br := bufio.NewReader(file)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			fmt.Println("来自缓存", string(line))
			resultstring = append(resultstring, string(line))

		}
		file.Close()

	} else {
		//不存在
		fmt.Println("文件bu存在")
		arrch := seg.CutForSearch("软件工程", true)
		var arr []string
		arr = append(arr, searchstr)
		arr = append(arr, returnstr(arrch)...)

		filepath := "Z:\\J\\洗币\\社会工程学\\newQQ\\Allqundata.txt"
		file, _ := os.Open(filepath) //打开文件
		br := bufio.NewReader(file)

		savefile, _ := os.Create(filelastpath)
		save := bufio.NewWriter(savefile)

		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}

			for j := 0; j < len(arr); j++ {
				if strings.Contains(string(line), arr[j]) {
					fmt.Println("来自搜索", string(line))
					fmt.Fprintln(save, string(line))
					//写入
					break //软件
				}
			}

		}
		save.Flush() //写入
		file.Close()
		savefile.Close()

	}

	return resultstring

}
func main() {
	mainSearch("软件工程")
	mainSearch("软件工程")
	mainSearch("软件工程")

}

func returnstr(ch <-chan string) []string {
	var strarr []string
	for word := range ch {
		fmt.Println(word)
		strarr = append(strarr, word)
	}
	return strarr
}
func main111x() {
	fmt.Println("文件bu存在")

	fmt.Println("xxx", returnstr(seg.CutForSearch("软件工程", true)))

}
