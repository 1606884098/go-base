package main

import (
	"fmt"
	"net"
	"os"
)

/*
关于聊天软件：
飞秋是根据是固定ip，局域网可可用，原来就是两个客户端tcp通讯
qq 微信不一样通过服务器连接。然后通过服务器中转  比如可以将qq号与连接对应  在服务器中转
11-conn
22-conn
首先用一个协程将上线的好放入map 管理连接
当 11 和22都上线了
11发消息给22时   11发送消息到服务器  服务器读取到消息 找到22-conn的连接，然后将消息写到22的conn里发送到22
*/
func main() {
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
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	// 主协程，接收服务器回发数据，打印至屏幕
	buf := make([]byte, 1024) // 定义用于存储服务器回发数据的切片缓冲区
	for {
		n, err := conn.Read(buf) // 从通信 socket 中读数据，存入切片缓冲区
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
	}

}
