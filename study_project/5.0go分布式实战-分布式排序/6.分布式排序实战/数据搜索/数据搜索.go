package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//N=86907937 行
func mainSearch(searchstr string) []string {

	filesavepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day6\\tmp\\"
	filelastpath := filesavepath + searchstr + ".txt"
	_, err := os.Stat(filelastpath)
	resultstring := []string{}
	if err == nil {
		//存在

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
			if strings.Contains(string(line), searchstr) {
				fmt.Println("来自搜索", string(line))
				fmt.Fprintln(save, string(line))
				//写入
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
