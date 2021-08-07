package main

import (
	"fmt"
	"go-base/study_project/go_web/chatroom/server/models"
	"net"
	"time"
)

func initUserDao() {
	models.MyUserDao = models.NewUserDao(pool)
}

func init() {
	//服务端一旦启动，初始化redis连接池
	initPool("localhost:6379", 16, 10, 300*time.Second)
	initUserDao()
}

func main() {
	fmt.Println("服务器在8080端口监听...")
	//服务器监听8080端口
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept err = ", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	//延迟关闭连接
	defer conn.Close()

	//创建总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯错误, err = ", err)
		return
	}
}
