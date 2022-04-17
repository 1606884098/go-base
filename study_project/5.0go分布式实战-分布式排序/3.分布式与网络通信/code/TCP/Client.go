package main

import (
	"fmt"
	"net"
)

//输入命令1ipconfig 1tasklist
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端连接失败")
		return
	}
	//接收服务器返回信息
	go func() {

		for {
			buf := make([]byte, 4096)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println("数据读取失败")
				return
			}
			fmt.Println("收到服务器信息长度", cnt, "内容", string(buf))
		}

	}()
	//等待给服务器输入信息
	for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input))
	}

}
