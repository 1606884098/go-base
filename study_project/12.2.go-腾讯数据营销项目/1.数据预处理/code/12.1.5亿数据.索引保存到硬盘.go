package main

import "os"
import "io"
import "bufio"
import (
	"fmt"

	"encoding/binary"
)

func makeIndexs() []int {
	N := 156774123
	indexdata := make([]int, N, N)
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodSort.txt"

	file, _ := os.Open(path) //打开文件
	br := bufio.NewReader(file)
	indexdata[0] = 0
	i := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		linestr := string(line)
		if i < N {
			indexdata[i] = len(linestr) + 1
		}
		i++
		if i%1000000 == 0 {
			fmt.Println(i)
		}

	}
	file.Close()
	fmt.Println("第一步读取长度索引", i)

	for j := 0; j < len(indexdata)-1; j++ {
		indexdata[j+1] += indexdata[j]
	}
	fmt.Println("第二部数组索引叠加完成")

	return indexdata

}

func main() {

	lengthlist := makeIndexs()
	fmt.Println("生成索引")
	savepath := "Z:\\J\\洗币\\社会工程学\\NBdata\\QQBigGoodindex.txt"
	savefile, _ := os.Create(savepath)
	wr := bufio.NewWriter(savefile)
	defer savefile.Close()
	for i := 0; i < len(lengthlist); i++ {
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(lengthlist[i]))
		//savefile.Write(b)
		wr.Write(b)
		if i%1000000 == 0 {
			fmt.Println(i)
		}
	}
	wr.Flush()

}
