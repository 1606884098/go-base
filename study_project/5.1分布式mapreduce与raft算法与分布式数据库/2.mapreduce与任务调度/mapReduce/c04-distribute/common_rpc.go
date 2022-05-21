package c04_distribute

import "net/rpc"

// 定义一个worker shutdown的响应
type ShutdownReply struct {
	Ntasks int // 代表指定worker执行的当前为止的任务数量(任务编号 )
}

// 添加一个注册结构，代表worker向主服务器注册时传递的参数
type ReigsterArgs struct {
	Worker string // worker的地址
}

// 任务参数结构, 存储任务相关的信息
type DoTaskArgs struct {
	JobName    string
	File       string   // 输入文件，只对map任务有用
	Phase      jobPhase // 任务类型
	TaskNumber int      // 任务编号
	// 另一个类型的任务总数
	// map需要该参数计算中间结果文件的输出数量
	// reduce需要该参数获取收集中间结果文件数量
	NumOtherPhase int
}

// 实现RPC请求发送函数
/*
	args:
		srv:地址
		rpcname:服务方法
		args:传递的参数
		reply:响应
*/
func call(srv string, rpcname string, args interface{}, reply interface{}) bool {

	// 连接rpc服务
	c, err := rpc.Dial("unix", srv)
	if err != nil {
		return false
	}
	defer c.Close()
	// 调用指定方法
	err = c.Call(rpcname, args, reply)
	if nil == err {
		return true
	}

	return false
}
