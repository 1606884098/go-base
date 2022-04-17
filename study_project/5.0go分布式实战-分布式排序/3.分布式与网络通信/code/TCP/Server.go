package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
)

func GoCmdWithResult(cmdstr string) string {
	cmdstr = strings.TrimSpace(cmdstr)
	fmt.Println(len(cmdstr), "-----------", cmdstr)

	cmd := exec.Command(cmdstr)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "error1"
	}
	if err := cmd.Start(); err != nil {
		return "error2"
	}
	outbyte, err := ioutil.ReadAll(stdout) //读取所有的输出
	stdout.Close()                         //关闭
	if err := cmd.Wait(); err != nil {
		return "error3"
	}
	return string(outbyte) //返回结果

}

//buf[]4096
//0 -----聊天
//1 -----命令
func Server(conn net.Conn) {
	if conn == nil {
		fmt.Println("无效连接")
		return
	}
	//接收数据，处理

	for {
		buf := make([]byte, 4096)
		cnt, err := conn.Read(buf) //接收数据
		if cnt == 0 || err != nil {
			conn.Close()
			return
		}
		recv := string(buf)
		fmt.Println("服务器收到", recv)
		if recv[0] == '0' {
			fmt.Println("收到聊天", recv[1:])
			//---处理聊天
			var input string
			fmt.Scanln(&input)
			conn.Write([]byte("0" + input))
		} else if recv[0] == '1' { //客户端命令前要加1 如1tasklist,1ipconfig
			fmt.Println("收到命令", recv[1:])
			//---命令执行
			conn.Write([]byte("1" + GoCmdWithResult(recv[1:cnt]))) //返回命令结果
		}

	}
}

func main() {

	server, err := net.Listen("tcp", "127.0.0.1:8848")
	defer server.Close()
	if err != nil {
		fmt.Println("服务器开启失败")
		return
	}
	fmt.Println("正在开启服务器....")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("连接出错")
		}
		go Server(conn) //并发处理
	}

}
