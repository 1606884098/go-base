package main

import "fmt"
import "os"
import "bufio"
import "io"
import (
	"strings"
	//"github.com/hashicorp/vault/helper/password"
)

// 数据抽离----密码单独抽取到一个文件
//  1234 123  1234  123   123  123  abc   xyz  abc 1234 1234 1234 1234 1234
//排序
//123   123   123  123  1234  1234  1234 1234 1234 1234 1234  abc  abc  xyz
//统计次数
// 123 -4  1234-7  abc-2  xyz-1
//排序
// 1234-7   123 -4   abc-2  xyz-1

//冒泡，插入，选择，快速

func main() {
	path := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQBig.txt"
	passfile, err := os.Open(path) //打开文件
	defer passfile.Close()

	if err != nil {
		fmt.Println("密码文件打开失败")
		return
	}
	br := bufio.NewReader(passfile) //读取文件对象

	savefile, err := os.Create("C:\\Users\\Tsinghua-yincheng\\Desktop\\day1综合project练习\\QQpassword.txt")
	defer savefile.Close()
	save := bufio.NewWriter(savefile) //写入

	for {
		line, _, end := br.ReadLine() //每次读取一行
		if end == io.EOF {
			break //文件结束跳出死循环
		}
		linestr := string(line) //读取每一行
		//fmt.Println(linestr)
		lines := strings.Split(linestr, "----")
		if len(lines) == 2 {
			password := lines[1]
			//fmt.Println(password)
			fmt.Fprintln(save, password)

		}

	}
	save.Flush() //刷新
}
