package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	// 创建监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close() // 主协程结束时，关闭listener
	fmt.Println("服务器等待客户端建立连接...")

	/*
		Accept()函数的作用是等待客户端的链接，如果客户端没有链接，该方法会阻塞。如果有客户端链接，那么该方
		法返回一个Socket负责与客户端进行通信。所以，每来一个客户端，该方法就应该返回一个Socket与其通信
	*/
	// 等待客户端连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	defer conn.Close() // 使用结束，断开与客户端链接
	fmt.Println("客户端与服务器连接建立成功...")
	// 接收客户端数据
	buf := make([]byte, 1024) // 创建1024大小的缓冲区，用于read
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("服务器读到:", string(buf[:n])) // 读多少，打印多少。

	cmd := exec.Command(string(buf[:n])) //相当于是一个后面小程序，客户端数据ifconfig 返回ip地址 liunx用于运维系统
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("收到" + string(out)))
}
