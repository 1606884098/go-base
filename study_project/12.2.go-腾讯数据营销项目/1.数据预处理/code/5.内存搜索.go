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
func LoadfiletoMEM() []string {
	const N = 86907937
	var arrlist []string = make([]string, N, N)
	filepath := "Z:\\J\\洗币\\社会工程学\\newQQ\\Allqundata.txt"
	file, _ := os.Open(filepath) //打开文件
	br := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		arrlist[i] = string(line)
		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println(i)
	return arrlist

}
func returnstr(ch <-chan string) []string {
	var strarr []string
	for word := range ch {
		fmt.Println(word)
		strarr = append(strarr, word)
	}
	return strarr
}
func main() {
	arrlist := LoadfiletoMEM()

	for {
		var input string
		fmt.Scanln(&input)
		fmt.Println("要搜索的是", input)

		filesavepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day7\\tmp\\"
		filelastpath := filesavepath + input + ".txt"
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

			savefile, _ := os.Create(filelastpath)
			save := bufio.NewWriter(savefile)

			arrch := seg.CutForSearch(input, true)
			var arr []string
			arr = append(arr, input)
			arr = append(arr, returnstr(arrch)...)

			for i := 0; i < len(arrlist); i++ {

				for j := 0; j < len(arr); j++ {
					if strings.Contains(string(arrlist[i]), arr[j]) {
						fmt.Println("来自搜索", string(arrlist[i]))
						//写入
						fmt.Fprintln(save, string(arrlist[i]))
						break //软件
					}
				}

			}
			save.Flush() //写入
			savefile.Close()

		}

	}

}
