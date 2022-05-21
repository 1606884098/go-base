package main

import (
	"fmt"
	"io/ioutil"
	//"github.com/pkg/errors"
)

//
//--
//----

func GetALLX(path string, level int) {
	fmt.Println("level", level)
	levelstr := ""
	if level == 1 {
		levelstr = "+"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "+"
	}

	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "\\" + fi.Name() //构造新的路径
			fmt.Println(levelstr + fulldir)
			//newlevel:=level+1
			//fmt.Println("call")
			GetALLX(fulldir, level+1) //文件夹递归处理

		} else {
			fulldir := path + "\\" + fi.Name() //构造新的路径
			fmt.Println(levelstr + fulldir)
		}

	}

}

func main() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1" //路径
	GetALLX(path, 1)                                      //抓取所有文件

}

func main2y() {
	//深度

}
