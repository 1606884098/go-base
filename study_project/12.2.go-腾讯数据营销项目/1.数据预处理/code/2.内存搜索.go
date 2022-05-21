package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadfiletoMEM() []string {
	const N = 156783175
	var arrlist []string = make([]string, N, N)
	filepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBig.txt"
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

func main() {
	arrlist := LoadfiletoMEM()

	for {
		var input string
		fmt.Scanln(&input)
		fmt.Println("要搜索的是", input)
		for i := 0; i < len(arrlist); i++ {
			if strings.Contains(arrlist[i], input) {
				fmt.Println("来自搜索", string(arrlist[i]))
				//写入
			}
		}
	}

}
