package main

import (
	"fmt"
	"net"
	"os"
)

func main2() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close() // 客户端终止时，关闭与服务器通信的 socket

	// 启动子协程，接收用户键盘输入
	go func() {
		str := make([]byte, 1024) // 创建用于存储用户键盘输入数据的切片缓冲区。
		for {                     // 反复读取
			n, err := os.Stdin.Read(str) // 获取用户键盘输入
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				return
			}
			// 将从键盘读到的数据，发送给服务器
			_, err = conn.Write(str[:n]) // 读多少，写多少
			if err != nil {
				fmt.Println("conn.Write err:", err)
				return
			}
		}
	}()

	// 主协程，接收服务器回发数据，打印至屏幕
	buf := make([]byte, 1024) // 定义用于存储服务器回发数据的切片缓冲区
	for {
		n, err := conn.Read(buf) // 从通信 socket 中读数据，存入切片缓冲区
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Printf("服务器回发：%s\n", string(buf[:n]))
	}

}
