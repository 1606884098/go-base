package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//遍历文件夹
func ForReadDir() {
	dir_list, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃")
	for i, v := range dir_list {
		fmt.Println(i, v.Name(), v.Size())
	}
}

//遍历二级文件夹
func ForReadDirSecond() {
	dir_list, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃")
	for _, v := range dir_list {
		dir_listx, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃" + "\\" + v.Name())
		for j, u := range dir_listx {
			filepath := "Z:\\J\\洗币\\社会工程学\\52G葫芦娃" + "\\" + v.Name() + "\\" + u.Name()
			fmt.Println(j, filepath)
		}
	}
}

//文件归并含义:新建一个文件----每个文件读取一次，每行写入。
func main() {
	savepath := "Z:\\J\\洗币\\社会工程学\\52G葫芦娃\\all163.txt"
	savefile, err := os.Create(savepath) //创建文件
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer savefile.Close()            //关闭
	save := bufio.NewWriter(savefile) //写入

	dir_list, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃")
	for _, v := range dir_list {
		dir_listx, _ := ioutil.ReadDir("Z:\\J\\洗币\\社会工程学\\52G葫芦娃" + "\\" + v.Name())
		for j, u := range dir_listx {
			filepath := "Z:\\J\\洗币\\社会工程学\\52G葫芦娃" + "\\" + v.Name() + "\\" + u.Name()
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
	}
	save.Flush()
}
