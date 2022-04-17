package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件归并含义。
//新建一个文件----每个文件读取一次，每行写入。

func main() {
	savepath := "Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\all163_3.6.txt"
	savefile, err := os.Create(savepath) //创建文件
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer savefile.Close()            //关闭
	save := bufio.NewWriter(savefile) //写入

	dir_listx, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\网易数据3.6G\\")
	for j, u := range dir_listx {
		filepath := "Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\网易数据3.6G\\" + u.Name()
		fmt.Println(filepath, "开始写入")
		fi, err := os.Open(filepath)
		if err != nil {
			fmt.Println("文件打开失败", filepath)
			continue
		}
		defer fi.Close()
		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break //文件到了末尾
			}
			fmt.Fprintln(save, string(line))
		}
		save.Flush()
		fmt.Println(filepath, "写入成功", j)

	}

	save.Flush()

}
