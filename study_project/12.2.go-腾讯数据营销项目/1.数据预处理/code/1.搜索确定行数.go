package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//N=86907937 行
func mainSearch(searchstr []string) {
	//不存在
	filepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBig.txt"
	file, _ := os.Open(filepath) //打开文件
	br := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		i++
		for j := 0; j < len(searchstr); j++ {
			if strings.Contains(string(line), searchstr[j]) {
				fmt.Println("来自搜索", string(line))
				//写入
			}
		}
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println(i)

}
func main() {
	const N = 156783175
	mainSearch([]string{"1652208024", "1196753634",
		"469470453", "568173044", "765642364", "250541506", "shangmeimei"})

}
