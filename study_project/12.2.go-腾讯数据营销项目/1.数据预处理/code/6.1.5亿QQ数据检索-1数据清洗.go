package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	fi, _ := os.Open("Z:J\\洗币\\社会工程学\\NBdata\\QQBig.txt")
	defer fi.Close()
	br := bufio.NewReader(fi)

	path := "Z:J\\洗币\\社会工程学\\NBdata\\QQBigGood.txt"
	savefile, _ := os.Create(path)
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //写入数据

	path1 := "Z:J\\洗币\\社会工程学\\NBdata\\QQBigBad.txt"
	savefile1, _ := os.Create(path1)
	defer savefile1.Close()
	save1 := bufio.NewWriter(savefile1) //写入数据

	reg1 := `^[1-9][0-9]{4,11}$`
	reg2 := `^.{6,16}$`
	rgx1 := regexp.MustCompile(reg1)
	rgx2 := regexp.MustCompile(reg2)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		lines := strings.Split(linestr, "----")
		if len(lines) == 2 {
			if rgx1.Match([]byte(lines[0])) && rgx2.Match([]byte(lines[1])) {
				fmt.Fprintln(save, lines[0]+"----"+lines[1])
			} else {
				fmt.Println("异常数据", linestr)
				fmt.Fprintln(save1, linestr)
			}
		} else {
			fmt.Println("异常数据", linestr)
			fmt.Fprintln(save1, linestr)
		}
	}
	save.Flush()
	save1.Flush()

}
