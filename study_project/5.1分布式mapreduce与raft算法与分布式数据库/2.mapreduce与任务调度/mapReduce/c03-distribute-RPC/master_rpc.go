package c03_distribute_RPC

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net"
	"net/rpc"
	"os"
)

// 开启RPC服务，等待worker进行注册
func (mr *Master) startRPCServer() {
	// 新建一个RPC实例对象
	rpcs := rpc.NewServer()
	// 注册master方法集
	// 只要满足RPC方法的规则，就可以是RPC方法
	rpcs.Register(mr)
	os.Remove(mr.address)
	// 监听
	l, err := net.Listen("unix", mr.address)
	if nil != err {
		log.Fatalf("RegisterServer %s error : %v!\n", mr.address, err)
	}
	// 将监听实例变成master的内部属性
	mr.l = l
	// 监听地址，获取连接
	go func() {
	loop:
		for {
			// 检测是否接收到中断
			select {
			case <-mr.shutdown:
				break loop
			default:

			}

			// 等待RPC连接
			conn, err := mr.l.Accept()
			if nil != err {
				log.Errorf("RegisterServer: accept error : %v\n", err)
				break
			} else {
				go func() {
					// 另起一个goroutine，运行rpc server
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			}
		}
		fmt.Println("RegisterServer: done!")

	}()
}

// 紧急中断
func (mr *Master) ShutDown(_, _ *struct{}) error {
	log.Errorf("Shutdown: registration server\n")
	close(mr.shutdown)
	mr.l.Close()
	return nil
}

// 正常停止master上的RPC服务
func (mr *Master) stopRPCServer() {
	var reply ShutdownReply
	// 调用一个发送RPC请求的函数
	ok := call(mr.address, "Master.Shudown",
		new(struct{}), &reply)
	if !ok {
		log.Errorf("RPC: Stop error!\n")
	}
	fmt.Println("stop registration donw")
}
