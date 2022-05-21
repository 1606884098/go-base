package main

import (
	"./Queue"
	"./StackArray"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func GetALL(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件加不可读取")
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "\\" + fi.Name() //构造新的路径
			files = append(files, fulldir)     //追加路径
			files, _ = GetALL(fulldir, files)  //文件夹递归处理

		} else {
			fulldir := path + "\\" + fi.Name() //构造新的路径
			files = append(files, fulldir)     //追加路径
		}

	}

	return files, nil
}

func main1x() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1" //路径
	files := []string{}                                   //数组字符串
	files, _ = GetALL(path, files)                        //抓取所有文件

	for i := 0; i < len(files); i++ { //打印路径
		fmt.Println(files[i])
	}

}
func main2x() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1" //路径
	files := []string{}                                   //数组字符串

	mystack := StackArray.NewStack()
	mystack.Push(path)

	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)     //加入列表
		read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "\\" + fi.Name() //构造新的路径
				//files=append(files,fulldir)//追加路径
				mystack.Push(fulldir)

			} else {
				fulldir := path + "\\" + fi.Name() //构造新的路径
				files = append(files, fulldir)     //追加路径
			}
		}

	}
	for i := 0; i < len(files); i++ { //打印
		fmt.Println(files[i])
	}

}
func main() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1" //路径
	files := []string{}                                   //数组字符串

	mystack := Queue.NewQueue()
	mystack.EnQueue(path)

	for {
		path := mystack.DeQueue() //不断从队列取出数据
		if path == nil {
			break
		}
		fmt.Println("get", path)
		read, _ := ioutil.ReadDir(path.(string)) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string) + "\\" + fi.Name() //构造新的路径
				//files=append(files,fulldir)//追加路径
				fmt.Println("Dir", fulldir) //文件夹

				mystack.EnQueue(fulldir)

			} else {
				fulldir := path.(string) + "\\" + fi.Name() //构造新的路径
				files = append(files, fulldir)              //追加路径
				fmt.Println("File", fulldir)                //文件夹
			}
		}

	}
	for i := 0; i < len(files); i++ { //打印
		fmt.Println(files[i])
	}
}
