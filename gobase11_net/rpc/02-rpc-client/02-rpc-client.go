package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

const HelloServiceName = "path/to/pkg.HelloService" //服务名

type HelloServiceInterface = interface { //接口
	Hello(request string, reply *string) error
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil) //匿名变量
//远程调用
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

//调用实际方法
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	/*	client, err := DialHelloService("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		var reply string
		err = client.Hello("TTTT", &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)*/

	//跨语言rpc
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}
	client1 := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply1 string
	err = client1.Call("HelloService.Hello", "我是跨语言", &reply1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply1)
}
