package c01_single_basic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"sort"
)

// 将多个reduce节点生成的最终结果输出文件输出到一个文件中进行排序，汇总
// 需要的参数
/*
	args:
		nRduce	节点数量
		jobName	任务名称
*/
func (mr *Master) merge(nReduce int, jobName string) {
	fmt.Println("Merge result...")
	kvs := make(map[string]string)
	for i := 0; i < nReduce; i++ {
		// 生成文件名
		p := mergeName(jobName, i)
		// 打开文件
		file, err := os.Open(p)
		if nil != err {
			log.Panicf("Merge : %v failed! %v\n", p, err)
		}

		decoder := json.NewDecoder(file)
		for decoder.More() {
			var kv KeyValue
			err := decoder.Decode(&kv)
			if nil != err {
				log.Errorf("Json decode failed! %v", err)
			}
			// 将具有相同key的内容进行合并
			kvs[kv.Key] = kv.Value
		}
		file.Close()
	}

	var keys []string
	for k := range kvs {
		keys = append(keys, k)
	}
	// 排序
	sort.Strings(keys)

	// 创建一个最终的输出结果文件
	file, err := os.Create("mrtmp." + jobName)
	if nil != err {
		log.Errorf("Merge: create failed! %v\n", err)
	}
	// 创建写入
	w := bufio.NewWriter(file)
	for _, k := range keys {
		fmt.Fprintf(w, "%s: %s\n", k, kvs[k])
	}
	w.Flush()
	file.Close()
}
