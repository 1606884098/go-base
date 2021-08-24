package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//创建客户端
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	//同步调用
	var req float32 //请求值
	req = 3
	var resp *float32 //返回值
	err = client.Call("MathUtil.CalculateCircleArea", req, &resp)
	//err = client.Call("bieming.CalculateCircleArea", req, &resp)//别名的方式调用
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)

	//异步调用
	/*	var respSync *float32
		//异步的调用方式
		syncCall := client.Go("bieming.CalculateCircleArea", req, &respSync, nil)
		replayDone := <-syncCall.Done
		fmt.Println(replayDone)
		fmt.Println(*respSync)*/
}
