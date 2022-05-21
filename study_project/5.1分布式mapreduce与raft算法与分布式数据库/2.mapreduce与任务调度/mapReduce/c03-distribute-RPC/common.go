package c03_distribute_RPC

import (
	"hash/fnv"
	"strconv"
)

// 自定义任务类型
type jobPhase string

const (
	mapPhase    jobPhase = "Map"
	reducePhase          = "Reduce"
)

// 用于保存需要传递给map和reduce函数的key/value数据对
type KeyValue struct {
	Key   string
	Value string
}

// reduce task输出文件名称
func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" +
		strconv.Itoa(reduceTask)
}

// 生成中间文件名称
func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) +
		"-" + strconv.Itoa(reduceTask)
}

// 哈希函数
func IHASH(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}
