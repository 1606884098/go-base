package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

//该方法向外暴露：提供计算圆形面积的服务
func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req //圆形的面积 s = π * r * r
	return nil                  //返回类型
}

//传递多个参数
type AddParma struct {
	Args1 float32 //第一个参数
	Args2 float32 //第二个参数
}

func (mu *MathUtil) Add(param AddParma, resp *float32) error {
	*resp = param.Args1 + param.Args2 //实现两数相加的功能
	return nil
}

func main() {
	mathUtil := new(MathUtil) //初始化指针数据类型
	//2、调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	//err := rpc.RegisterName("bieming", mathUtil)
	if err != nil {
		panic(err.Error())
	}
	//3、通过该函数把mathUtil中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()
	//4、在特定的端口进行监听
	listen, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic(err.Error())
	}

	http.Serve(listen, nil)

}
