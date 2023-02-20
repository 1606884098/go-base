package main

import (
	"fmt"
	"net"
)

func main() {

	// 主动发起连接请求相当于三次握手
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	// 结束时，关闭连接，相当于四次挥手
	defer conn.Close()

	for { //发送数据
		var input string
		fmt.Scanf("%s", &input)             //扫描输入
		n, err := conn.Write([]byte(input)) //向服务端发送数据  如ipconfig
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("写入成功：", n)

		buf := make([]byte, 4096)
		n, _ = conn.Read(buf) //从服务端接收数据
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("收到服务端的数据：", string(buf[:n]))
	}
}
