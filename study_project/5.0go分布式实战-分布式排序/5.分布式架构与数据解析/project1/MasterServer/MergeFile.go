package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	//"strings"
	"strconv"
	"strings"
)

type Pass struct {
	PassWord string
	times    int
}

//内置了归并排序
//归并算法
func Merge(arr []string, filepath string, datatype int, SmalltoBig bool, filename string) string {

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

			//filename:="csdnPasswd"
			//csdn1
			v3 := filename + getnumberstr(v1) + getnumberstr(v2) + "sort.txt"
			fmt.Println(v3)
			if !strings.Contains(v1, "\\") {
				v1 = filepath + v1
			}
			if !strings.Contains(v2, "\\") {
				v2 = filepath + v2
			}
			Mergre(v1, v2, filepath+v3, datatype, SmalltoBig)

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
func getnumberstr(str1 string) string {
	//str1:="163-1.txt"
	reg := regexp.MustCompile(`(\d+)sort.`)
	ss := reg.FindAllStringSubmatch(str1, -1)
	//fmt.Println(strconv.Itoa(int(ss[0][1]-48)))
	fmt.Println(str1)
	fmt.Println("ss---", ss, ss[0][1])
	return ss[0][1]
}

//解决两个文件归并排序，实现泛用
func Mergre(path1, path2, pathsave string, datatype int, SmalltoBig bool) {
	fmt.Println("path1", path1)
	fmt.Println("path2", path2)
	fmt.Println("pathsave", pathsave)

	var BigSmall func(data1, data2 interface{}) bool
	if datatype == 1 {
		if SmalltoBig {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(int) < data2.(int)
			}
		} else {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(int) > data2.(int)
			}
		}
	} else if datatype == 2 {
		if SmalltoBig {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(float64) < data2.(float64)
			}
		} else {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(float64) > data2.(float64)
			}
		}
	} else if datatype == 3 {
		if SmalltoBig {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(string) < data2.(string)
			}
		} else {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(string) > data2.(string)
			}
		}
	} else if datatype == 4 {
		if SmalltoBig {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(Pass).times < data2.(Pass).times
			}
		} else {
			BigSmall = func(data1, data2 interface{}) bool {
				return data1.(Pass).times > data2.(Pass).times
			}
		}

	}

	length1 := GetLineNumber(path1)
	length2 := GetLineNumber(path2) //文件的长度

	fmt.Println("length", length1, length2)

	fi1, _ := os.Open(path1)
	fi2, _ := os.Open(path2)
	defer fi1.Close()
	defer fi2.Close()
	br1 := bufio.NewReader(fi1)
	br2 := bufio.NewReader(fi2) //读取文件

	savefile, err := os.Create(pathsave)
	if err != nil {
		fmt.Println("pathsave", err)
	}
	save := bufio.NewWriter(savefile) //保存

	i, j := 0, 0
	line1, _, _ := br1.ReadLine()
	line2, _, _ := br2.ReadLine()

	for i < length1 && j < length2 {
		var in1, in2 interface{}
		if datatype == 1 {
			in1, _ = strconv.Atoi(string(line1))
			in2, _ = strconv.Atoi(string(line2))
		} else if datatype == 2 {
			in1, _ = strconv.ParseFloat(string(line1), 64)
			in2, _ = strconv.ParseFloat(string(line2), 64)
		} else if datatype == 3 {
			in1 = string(line1)
			in2 = string(line2)
		} else if datatype == 4 {
			line1list := strings.Split(string(line1), " # ")
			times, _ := strconv.Atoi(line1list[1])
			in1 = Pass{line1list[0], times}

			line2list := strings.Split(string(line2), " # ")
			times2, _ := strconv.Atoi(line2list[1])
			in2 = Pass{line2list[0], times2}
		}

		if BigSmall(in1, in2) {
			//fmt.Println(string(line2))
			fmt.Fprintln(save, string(line2))
			line2, _, _ = br2.ReadLine()
			j++

		} else if BigSmall(in2, in1) {
			//fmt.Println(string(line1))
			fmt.Fprintln(save, string(line1))
			line1, _, _ = br1.ReadLine()
			i++

		} else {
			//fmt.Println(string(line1))
			//fmt.Println(string(line2))
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
		//fmt.Println(string(line1))
		line1, _, _ = br1.ReadLine()
		i++
	}
	for j < length2 {
		fmt.Fprintln(save, string(line2))
		//fmt.Println(string(line2))
		line2, _, _ = br2.ReadLine()
		j++
	}
	save.Flush()
	savefile.Close()
}
