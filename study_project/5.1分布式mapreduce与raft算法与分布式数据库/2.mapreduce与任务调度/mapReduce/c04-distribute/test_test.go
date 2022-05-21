package c04_distribute

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	nNumber = 100
	nMap    = 10
	nReduce = 5
)

// 创建一个包含N个编号的输入文件
// 通过mapReduce进行处理
// 检查最终输出文件中是否包含了N个编号

// 自定义的map分割处理函数
func MapFunc(file string, value string) (res []KeyValue) {
	words := strings.Fields(value)
	for _, w := range words {
		kv := KeyValue{w, ""}
		res = append(res, kv)
	}
	return
}

// 自定义的reduce聚合函数
func ReduceFunc(key string, values []string) string {

	return ""
}

// 顺序执行的mapReduce
func TestSequentialSignle(t *testing.T) {
	Sequential("test",
		makeInputs(1),
		1,
		MapFunc, ReduceFunc)
}

// 顺序执行的mapReduce
// 中间文件数量nMap*nReduce
func TestSequentialMany(t *testing.T) {
	Sequential("test",
		makeInputs(5),
		3,
		MapFunc, ReduceFunc)
}

// 创建输入文件
// 根据指定的数量创建输入文件，返回创建好的文件名列表
// 写入相应的数据
// count : 创建的文件数量
func makeInputs(num int) []string {
	var names []string
	var i = 0
	for f := 0; f < num; f++ {
		// 文件命名方式 ： 根据 mit6.824课程命名
		names = append(names, fmt.Sprintf("824-mrinput-%d.txt", f))
		// 创建文件
		file, err := os.Create(names[f])
		if nil != err {
			log.Fatalf("create input file [%s] failed. error:", file, err)
		}
		w := bufio.NewWriter(file)
		for i < (f+1)*(nNumber/num) {
			// 写入i到w中
			fmt.Fprintf(w, "%d\n", i)
			i++
		}
		// 把buffer中的内容写入文件
		w.Flush()
		file.Close()
	}

	return names
}

func setup() *Master {
	fmt.Println("SETUP Master")
	files := makeInputs(nMap)
	master := "master"
	mr := Distributed("test", files, nReduce, master)
	return mr
}

// 设置worker标识
func workerFlag(num int) string {
	s := "824-"
	s += strconv.Itoa(os.Getuid()) + "/"
	os.Mkdir(s, 0777)
	s += "mr"
	s += strconv.Itoa(os.Getpid()) + "-" + strconv.Itoa(num)
	return s
}

func TestBasic(t *testing.T) {
	// 启动master
	mr := setup()
	// 启动worker
	for i := 0; i < 2; i++ {
		go RunWorker(mr.address, workerFlag(i), MapFunc, ReduceFunc, -1)
	}
}
