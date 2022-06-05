package main

import (
	"fmt"
	"net"
	"os/exec"
)

func MsgHandler(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {

		}
		//clientip:=conn.RemoteAddr()//远程地址
		if n != 0 {
			if string(buf[0:1]) == "0" {
				fmt.Println("client  data", string(buf[1:n]))

				conn.Write([]byte("收到数据:" + string(buf[1:n]) + "\n"))
			} else {
				fmt.Println("client  cmd", string(buf[1:n]))
				cmd := exec.Command(string(buf[1:n])) //执行命令
				cmd.Run()
				conn.Write([]byte("收到命令:" + string(buf[1:n]) + "\n"))
			}

		}

	}

}
func main() {
	server_listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		panic(err) //处理错误
	}
	defer server_listener.Close() //延迟关闭
	for {
		new_conn, err := server_listener.Accept() //接收消息
		if err != nil {
			panic(err) //处理错误
		}
		go MsgHandler(new_conn) //处理客户端消息

	}

}
