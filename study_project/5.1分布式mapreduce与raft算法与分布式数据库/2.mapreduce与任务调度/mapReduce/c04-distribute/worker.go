package c04_distribute

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

// 工作者，等待DoTask任务或者ShutDown任务
type Worker struct {
	sync.Mutex
	name       string // worker名称
	Map        func(string, string) []KeyValue
	Reduce     func(string, []string) string
	nTasks     int // 执行的任务问题
	concurrent int // 当前worker并发执行任务量
	l          net.Listener

	nRPC int // 退出标志
}

// 任务执行函数，在有新的任务被分配给该workers时，master调用该函数执行任务
func (wk *Worker) DoTask(args *DoTaskArgs, _ *struct{}) error {
	// 任务数量加1
	wk.Lock()
	wk.nTasks += 1
	wk.Unlock()
	switch args.Phase {
	case mapPhase:
		doMap(args.JobName, args.TaskNumber, args.File, args.NumOtherPhase, wk.Map)
	case reducePhase:
		doReduce(args.JobName, args.TaskNumber,
			mergeName(args.JobName, args.TaskNumber), args.NumOtherPhase, wk.Reduce)
	}
	fmt.Printf("%s:%v task #%d done\n", wk.name, args.Phase, args.TaskNumber)
	return nil
}

// 启动worker，与master建立连接
func RunWorker(MasterAddress string, me string,
	MapFunc func(string, string) []KeyValue,
	ReduceFunc func(string, []string) string,
	nRPC int) {
	fmt.Printf("RunWorker:%s\n", me)
	wk := new(Worker)
	wk.name = me
	wk.Map = MapFunc
	wk.Reduce = ReduceFunc
	wk.nRPC = nRPC
	// 新建一个rpc服务实例
	rpcs := rpc.NewServer()
	// 注册
	rpcs.Register(wk)
	os.Remove(me)

	l, err := net.Listen("unix", me)
	if nil != err {
		log.Fatalf("RunWorker: Worker %s error %v\n", me, err)
	}
	wk.l = l
	// 注册到master中
	wk.register(MasterAddress)

	for {
		wk.Lock()
		// 没有连上
		if wk.nRPC == 0 {
			wk.Unlock()
			break
		}
		conn, err := wk.l.Accept()
		if nil != err {
			break
		} else {
			wk.Lock()
			wk.nRPC--
			wk.Unlock()
			go rpcs.ServeConn(conn)
		}
		wk.l.Close()
		fmt.Printf("RunWorker %s exit\n", me)
	}
}

// 注册，相当于告知master, worker的存在
func (wk *Worker) register(master string) {
	args := new(ReigsterArgs)
	args.Worker = wk.name
	// 调用master的注册函数，注册worker
	ok := call(master, "Master.Register", args, new(struct{}))
	if !ok {
		log.Errorf("Register: RPC %s master error\n", master)
	}
}

//
func (wk *Worker) Shutdown(_ *struct{}, res *ShutdownReply) error {
	fmt.Printf("Shutdown %s\n", wk.name)
	wk.Lock()
	defer wk.Unlock()
	res.Ntasks = wk.nTasks
	wk.nRPC = 1
	return nil
}
