package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //初始化
)

func main() {
	dsn := "root:ye668899@tcp(localhost:3306)/grpc-todo-list"
	//打开数据库链接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db err:", err)
		return
	}
	//关闭数据库链接
	defer db.Close()
	fmt.Println("数据库链接成功")

}
