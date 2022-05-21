package c04_distribute

import (
	"fmt"
	"sync"
)

// 调度函数的实现，决定如何向worker分配任务
func schedule(jobName string,
	mapFiles []string,
	nReduce int,
	phase jobPhase,
	registerChan chan string,
) {
	// 当前任务数量
	var ntasks int
	// 另一个任务数量
	var n_other int
	// 类型判断
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}
	var lock *sync.Mutex = &sync.Mutex{}
	// 执行mapReduce的函数功能调用操作
	// 所有任务必须调度给workers，在等待所有任务全部执行成功之后，再下一步处理
	// 一个worker可以完成多个任务
	var wg sync.WaitGroup
	// 声明一个任务列表，将所有待处理的任务添加进去
	tasks := make([]int, ntasks)
	for i := 0; i < ntasks; i++ {
		tasks[i] = i
	}

	for {
		lock.Lock()
		// 没有任务需要执行
		if len(tasks) <= 0 {
			lock.Unlock()
			break
		}
		// 执行任务
		// 每执行一个任务，都将该任务从任务列表中删除
		task := tasks[0]
		tasks = append(tasks[:0], tasks[1:]...)
		lock.Unlock()
		// 任务参数赋值
		var doTaskArgs *DoTaskArgs = &DoTaskArgs{
			JobName:       jobName,
			Phase:         phase,
			TaskNumber:    task,
			NumOtherPhase: n_other,
		}
		if phase == mapPhase {
			doTaskArgs.File = mapFiles[task]
		}

		worker := <-registerChan
		wg.Add(1)
		// 调度RPC
		go func() {
			ok := call(worker, "Worker.DoTask", doTaskArgs, nil)
			if ok {
				wg.Done()
			} else {
				// worker执行任务失败
				// 将该任务重新加入列表，相当于重新进行任务接入
				lock.Lock()
				tasks = append(tasks, doTaskArgs.TaskNumber)
				lock.Unlock()
				wg.Done()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Schedule: %v phase done\n", phase)
}
