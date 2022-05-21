package c04_distribute

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net"
	"sync"
)

// master结构
type Master struct {
	// master地址
	address string
	// 锁
	sync.Mutex
	// 存储worker的缓存,存储套接字，代表RPC地址
	workers  []string
	jobName  string   // 当前要执行的任务名称
	files    []string // 输入文件
	nReduce  int      // reduce分区数量
	newCond  *sync.Cond
	l        net.Listener  // 监听对象
	shutdown chan struct{} // 中断服务
	stats    []int
}

//  初始化master实例
func newMaster(master string) *Master {
	mr := new(Master)
	mr.address = master
	mr.shutdown = make(chan struct{})
	mr.newCond = sync.NewCond(mr)
	return mr
}

// 任务调度函数(顺序执行)
func Sequential(
	jobName string, // 任务名称
	files []string, // 输入文件(待处理文件)
	nReduce int, // 分区数量
	mapF func(string, string) []KeyValue,
	reduceF func(string, []string) string) {
	/*....*/

	// 执行分配的任务
	// 在mapReduce中。任务主要分为 map任务和redcue任务
	mr := newMaster("master")
	mr.run(jobName, files, nReduce, func(phase jobPhase) {
		switch phase {
		case mapPhase:
			//执行map任务
			// map任务的调用次数由输入文件的个数决定
			for i, f := range files {
				// doMap
				doMap(mr.address, i, f, mr.nReduce, mapF)
			}
		case reducePhase:
			//执行reduce任务
			// reduce任务的调用次数由nReduce大小来决定决定
			for i := 0; i < mr.nReduce; i++ {
				doReduce(mr.jobName, i, mergeName(mr.jobName, i), len(files), reduceF)
			}
		}
	}, nil) // 实际执行
}

// 实际上的执行函数
// 执行给定的任务
func (mr *Master) run(
	jobName string,
	files []string,
	nreduce int,
	schedule func(phase jobPhase),
	finish func(), // 完成此次任务之后，需要执行的逻辑，收尾
) {
	// 执行Map任务
	schedule(mapPhase)
	// 执行Reduce任务
	schedule(reducePhase)
	mr.jobName = jobName
	mr.files = files
	mr.nReduce = nreduce
	finish()
	// 合并
	mr.merge()
}

// 实现一个worker注册函数，这是一个RPC的方法
func (mr *Master) Register(args *ReigsterArgs, _ *struct{}) error {
	// 加锁
	mr.Lock()
	defer mr.Unlock()
	// 注册worker
	mr.workers = append(mr.workers, args.Worker)
	// 广播给其它节点有新的worker进入
	mr.newCond.Broadcast()
	return nil
}

// 实现一个worker传递函数，将所有的已经存在的workers与新注册的worker传递到一个
// 通道中，让调度函数进行接收处理
func (mr *Master) forwardRegistraions(ch chan string) {
	i := 0
	for {
		mr.Lock()
		if len(mr.workers) > i {

			w := mr.workers[i]
			go func() { ch <- w }()
			i = i + 1
		} else {
			mr.newCond.Wait()
		}
		mr.Unlock()
	}
}

// 分布执行mapReduce任务
// 向通过RPC在主服务器注册的workers调度map和reduce任务
func Distributed(jobName string,
	files []string,
	nreduce int,
	master string,
) (mr *Master) {
	fmt.Println("Distributed...")
	mr = newMaster(master)
	// 启动RPC服务
	mr.startRPCServer()
	// 执行一个任务
	go mr.run(jobName, files, nreduce, func(phase jobPhase) {
		ch := make(chan string)
		go mr.forwardRegistraions(ch)
		// 调度执行
		schedule(mr.jobName, mr.files, mr.nReduce, phase, ch)
	}, func() {
		mr.stats = mr.killWorker()
		mr.stopRPCServer()
	})
	return
}

// 实现一个清理 worker的函数
// 该函数通过向每个worker发送shutDown RPC请求来清理worker
// 返回workers中，worker已经执行的任务数量
func (mr *Master) killWorker() []int {
	mr.Lock()
	defer mr.Unlock()
	ntasks := make([]int, 0, len(mr.workers))
	for _, w := range mr.workers {
		fmt.Printf("Master:shutdown worker %s\n", w)
		var reply ShutdownReply
		// 调用worker的shutdown方法
		ok := call(w, "Worker.Shutdown", new(struct{}), &reply)
		if !ok {
			log.Errorf("Master: RPC %s shutdown failed!\n", w)
		} else {
			ntasks = append(ntasks, reply.Ntasks)
		}
	}
	return ntasks
}
