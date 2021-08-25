package main

import (
	"fmt"
	"net/rpc"
)

//传递多个参数  理论上来讲这个是DTO，可以单独一个包，然后client server端都导入
type AddParmaTwo struct {
	Args1 float32 //第一个参数
	Args2 float32 //第二个参数
}

func main() {
	//创建客户端
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	//同步调用
	var req float32 = 3 //请求值
	//req = 3
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
	var addParma AddParmaTwo
	addParma.Args1 = 10
	addParma.Args2 = 12
	err = client.Call("MathUtil.Add", addParma, &resp)
	//err = client.Call("bieming.CalculateCircleArea", req, &resp)//别名的方式调用
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)
}
