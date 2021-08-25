package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService" //服务名

type HelloServiceInterface = interface { //接口
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error { //操作接口
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
	//跨语言rpc
	/*	rpc.RegisterName("HelloService", new(HelloService))
		listener, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal("ListenTCP error:", err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}
			go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		}*/

}

type clientRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params"`
	Id     uint64         `json:"id"`
}
