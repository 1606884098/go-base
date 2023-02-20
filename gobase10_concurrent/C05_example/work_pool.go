package main

/*
使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，
计算其每个位数上的数字之和为19。
1.开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
2.开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3.主 goroutine 从resultChan取出结果并打印到终端输出
*/
import (
	"fmt"
	"math/rand"
	"sync"
)

var swg sync.WaitGroup

type Job struct {
	Id      int
	RandNum int
}
type Result struct {
	job *Job
	sum int
}

func main() {

	// chan中是要存数据的，整型、字符串、布尔、结构体、只够提指针
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	// 开启一个goroutine从resultChan中取数据打印（消费数据）
	/* 一直处理使用此方法
	   go func(result <-chan *Result) {
	       for ret := range result {
	           fmt.Printf("id:%d, randnum:%d, result:%d\n", ret.job.Id, ret.job.RandNum, ret.sum)
	       }
	       swg.Done()
	   }(resultChan)
	*/
	//
	swg.Add(1)
	go func(result <-chan *Result) {
		for i := 0; i < 10; i++ {
			ret := <-result
			fmt.Printf("id: %d, randNum: %d, sum: %d\n", ret.job.Id, ret.job.RandNum, ret.sum)
		}
		close(resultChan)
		swg.Done()
	}(resultChan)

	// 开启goroutine池去从jobChan中取数据，处理完后在发送到resultChan中
	createPool(64, jobChan, resultChan)

	// 主goroutine负责生产随机数并发往通道（生产数据）
	var i int
	for i < 10 {
		i++
		randNum := rand.Int()
		job := &Job{
			Id:      i,
			RandNum: randNum,
		}
		// 结构体发送到通道也是拷贝，如果不想拷贝，可以发送结构体指针到通道中
		jobChan <- job
	}
	close(jobChan)

	swg.Wait()

}

// 创建工作池，参数num:开启几个协程
func createPool(num int, jobChan <-chan *Job, resultChan chan<- *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan <-chan *Job, resultChan chan<- *Result) {
			for job := range jobChan {
				r_num := job.RandNum // 接收随机值
				var sum int          // 存取结果值
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r // 发送到结果通道
			}
		}(jobChan, resultChan)
	}
}
