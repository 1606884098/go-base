package c04_distribute

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

// 实现一个map任务管理函数，从input files中读取内容
// 将输出分成指定数量的中间文件
// 自定义分割标准
func doMap(
	jobName string, // 任务名称
	mapTaskNumber int, // 当前map任务编号
	inFile string, // 输入文件
	nReduce int, // 当前map任务执行的reduce编号
	mapF func(file string, contents string) []KeyValue) {
	// 打开指定文件
	f, err := os.Open(inFile)
	if nil != err {
		log.Errorf("open file %s failed! %v\n", inFile, err)
	}
	defer f.Close()
	// 从输入文件inFile中读取内容
	content, err := ioutil.ReadAll(f)
	if nil != err {
		log.Errorf("read the content of file failed! %v\n", err)
	}
	// 通过调用mapF对内容进行处理，分割map任务输出
	// 将指定文件的内容解析为key-value
	kvs := mapF(inFile, string(content))
	// 生成一个编码对象
	// 每一个map任务生成nReduce个中间文件对象
	encoders := make([]*json.Encoder, nReduce)
	// 创建nReduce个中间结果文件
	for i := 0; i < nReduce; i++ {
		// 生成中间文件的名称
		file_name := reduceName(jobName, mapTaskNumber, i)
		f, err := os.Create(file_name)
		if nil != err {
			log.Errorf("unable to create file [%s]: %v\n", file_name, err)
		}
		defer f.Close()
		encoders[i] = json.NewEncoder(f)
	}

	// 将kvs中的内容存入前面生成的中间文件中去
	for _, v := range kvs {
		// 自定义规则对key值进行分类
		// 此处以编号的哈希值对nReduce取余进行分类
		index := IHASH(v.Key) % nReduce
		if err := encoders[index].Encode(&v); nil != err {
			log.Errorf("Unable to write file\n")
		}
	}
}
