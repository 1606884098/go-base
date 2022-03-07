package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	//"strconv"
)

func GetLineNumber(filepath string) int {
	fi, err := os.Open(filepath)
	defer fi.Close() //打开文件
	if err != nil {
		return -1
	}
	i := 0 //统计行数
	br := bufio.NewReader(fi)
	for {
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break //文件结束，末尾
		}
		i++ //统计行数
	}
	return i

}
func Mergre(path1, path2, pathsave string) {
	length1 := GetLineNumber(path1)
	length2 := GetLineNumber(path2) //文件的长度
	fi1, _ := os.Open(path1)
	fi2, _ := os.Open(path2)
	defer fi1.Close()
	defer fi2.Close()
	br1 := bufio.NewReader(fi1)
	br2 := bufio.NewReader(fi2) //读取文件

	savefile, _ := os.Create(pathsave)
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //保存
	i, j := 0, 0
	line1, _, _ := br1.ReadLine()
	line2, _, _ := br2.ReadLine()

	for i < length1 && j < length2 {
		if string(line1) < string(line2) {
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++

		} else if string(line1) > string(line2) {
			fmt.Fprintln(save, string(line1))
			line1, _, _ = br1.ReadLine()
			i++

		} else {
			fmt.Fprintln(save, string(line1))
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++
			line1, _, _ = br1.ReadLine()
			i++
		}
	}
	for i < length1 {
		fmt.Fprintln(save, string(line1))
		line1, _, _ = br1.ReadLine()
		i++
	}
	for j < length2 {
		fmt.Fprintln(save, string(line2))
		line2, _, _ = br2.ReadLine()
		j++
	}
	save.Flush()
}

func Merge(arr []string, filepath string) string {
	mylist := list.New()
	for i := 0; i < len(arr); i++ {
		mylist.PushBack(arr[i]) //数据批量压入
	}
	fmt.Println(mylist.Len()) //栈数据长度

	for mylist.Len() != 1 {
		e1 := mylist.Back()
		mylist.Remove(e1)

		e2 := mylist.Back()
		mylist.Remove(e2) //取得两个数据

		if e1 != nil && e2 != nil { //两个数据不为空，归并
			v1, _ := e1.Value.(string)
			v2, _ := e2.Value.(string)
			v3 := "163-" + getnumberstr(v1) + getnumberstr(v2) + ".txt"
			fmt.Println(v3)
			Mergre(filepath+v1, filepath+v2, filepath+v3)
			mylist.PushBack(v3)
		} else if e1 != nil && e2 == nil { //一个不为空，另外一个为空，再次压入
			v1, _ := e1.Value.(string)
			mylist.PushBack(v1)
		} else if e1 == nil && e2 == nil { //均为空，跳出循环
			break
		} else {
			break
		}

	}

	return mylist.Back().Value.(string)

}

func main() {
	dir_list, _ := ioutil.ReadDir("C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\datasort\\")
	filenames := []string{}
	filepaths := []string{}
	filepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\datasort\\"
	for _, v := range dir_list {
		filenames = append(filenames, v.Name())
		filepaths = append(filepaths, filepath+v.Name())
		fmt.Println(v.Name())
		fmt.Println(filepath + v.Name())

	}
	fmt.Println(Merge(filenames, filepath))
}

// 1  2  3  4  5  6  7
// 12  34  56  7
//  1234  567
//  1234567

func getnumberstr(str1 string) string {
	//str1:="163-1.txt"
	reg := regexp.MustCompile(`-(\d+).`)
	ss := reg.FindAllStringSubmatch(str1, -1)
	//fmt.Println(strconv.Itoa(int(ss[0][1]-48)))
	return ss[0][1]
}

func main1x() {
	str1 := "163-1234.txt"
	reg := regexp.MustCompile(`-(\d+).`)
	ss := reg.FindAllStringSubmatch(str1, -1)
	fmt.Println(ss[0][1])
	//fmt.Println(strconv.Itoa(int(ss[0][1]-48)))

}
