package pipelineMiddleWare

import (
	"bufio"
	"net"
)

//给我一个ip端口 ，127.0.0.1:8090 //写入数据
func NetWordkWrite(addr string, in <-chan int) {
	listen, err := net.Listen("tcp", addr) //监听
	if err != nil {
		panic(err)
	}
	go func() {
		defer listen.Close() //关闭网络

		conn, err := listen.Accept() //接收信息
		if err != nil {
			panic(err)
		}
		defer conn.Close() //关闭链接

		writer := bufio.NewWriter(conn) //写入数据
		defer writer.Flush()
		WriterSlink(writer, in)

	}()

}

//给我一个端口，读取数据
func NetWordkRead(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v //压入数据
		}

		close(out)
	}()

	return out
}
