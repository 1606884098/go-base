package main

import "fmt"
import "os"

func main() {

	findex, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\qq_data\\index.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer findex.Close() //延迟关闭文件
	fii, err := os.Open("Z:\\E\\区块链项目\\洗币\\社会工程学\\newQQ\\qq_data\\qq_name_qun.txt")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fii.Close() //延迟关闭文件

	for {
		var input int
		fmt.Scanf("%d", &input)
		mid := input - 1
		findex.Seek(int64(mid*15), 0)
		bx := make([]byte, 15, 15)
		_, _ = findex.Read(bx)
		var data int64
		fmt.Sscanf(string(bx), "%15d", &data) //数组载入内存

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
		fmt.Println(midstr)

	}

}
