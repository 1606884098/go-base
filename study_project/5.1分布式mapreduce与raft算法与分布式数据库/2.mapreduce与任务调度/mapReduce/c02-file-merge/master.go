package c01_single_basic

// master结构
type Master struct {
}

// 任务调度函数
func Sequential(
	jobName string, // 任务名称
	files []string, // 输入文件(待处理文件)
	nReduce int, // 分区数量
	mapF func(string, string) []KeyValue,
	reduceF func(string, []string) string) {
	/*....*/

	// 执行分配的任务
	// 在mapReduce中。任务主要分为 map任务和redcue任务
	mr := newMaster()
	mr.run(jobName, files, nReduce, func(phase jobPhase) {
		switch phase {
		case mapPhase:
			//执行map任务
			// map任务的调用次数由输入文件的个数决定
			for i, f := range files {
				// doMap
				doMap(jobName, i, f, nReduce, mapF)
			}
		case reducePhase:
			//执行reduce任务
			// reduce任务的调用次数由nReduce大小来决定决定
			for i := 0; i < nReduce; i++ {
				doReduce(jobName, i, mergeName(jobName, i), len(files), reduceF)
			}
		}
	}) // 实际执行
}

//  初始化master实例
func newMaster() *Master {

	return nil
}

// 实际上的执行函数
// 执行给定的任务
func (mr *Master) run(jobName string, files []string, nreduce int,
	schedule func(phase jobPhase)) {
	// 执行Map任务
	schedule(mapPhase)
	// 执行Reduce任务
	schedule(reducePhase)
	// 合并
	mr.merge(nreduce, jobName)
}
