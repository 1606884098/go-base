package c03_distribute_RPC

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
)

// 管理reduce任务
func doReduce(
	jobName string, // 任务名称
	reduceTaskNumber int, // reduce任务编号
	outFile string, // 输出文件
	nMap int, // 运行的map任务号
	reduceF func(key string, value []string) string) {
	var result map[string][]string = make(map[string][]string)
	// 打开每一个中间文件
	for i := 0; i < nMap; i++ {
		interFile := reduceName(jobName, i, reduceTaskNumber)
		f, err := os.Open(interFile)
		if nil != err {
			log.Errorf("read content from file [%s] failed! %v\n", interFile, err)
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		var kv KeyValue
		for decoder.More() {
			err := decoder.Decode(&kv)
			if nil != err {
				log.Errorf("Json decode failed! %v", err)
			}
			// 将具有相同key的内容进行合并
			result[kv.Key] = append(result[kv.Key], kv.Value)
		}
	}

	// 把内容做相应的处理，然后存入最终输出文件
	var keys []string
	for key, _ := range result {
		keys = append(keys, key)
	}
	// 新建输出文件
	out_file, err := os.Create(outFile)
	if nil != err {
		log.Errorf("create outFile faile! %v", err)
	}
	defer out_file.Close()

	encoder := json.NewEncoder(out_file)
	for _, key := range keys {
		encoder.Encode(KeyValue{key, reduceF(key, result[key])})
	}
}
