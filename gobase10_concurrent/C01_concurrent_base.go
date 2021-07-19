package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese" //go get -v golang.org/x/text/encoding/simplifiedchinese
	"golang.org/x/text/transform"                  //需要go get golang.org/x/text/transform
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	//1.windows进程使用
	create_windows_process() //创建进程
	Stdout_do()
	pipe_mathod()     //windows管道的使用
	process_message() //关于进程信息
	//2.linux进程使用
	create_linux_process()
}

func create_linux_process() {
	cmd := exec.Command("ls", "-lah")
	fmt.Println(cmd.Process.Pid) //进程的pid
	cmd.Stdout = os.Stdout       //system output
	cmd.Stderr = os.Stderr       //system err
	cmd.Run()

	path, err := exec.LookPath("go") //查看程序是否存在
	if err != nil {
		fmt.Println("no find")
	} else {
		fmt.Println(path)
	}
}

func Stdout_do() {
	//关于stdout的用法
	cmd_1 := exec.Command("ping", "www.baidu.com")
	var stdout, stderr bytes.Buffer //创建二进制输入，区别输出，区别错误
	cmd_1.Stdout = &stdout          //设定输出错误，输出
	cmd_1.Stderr = &stderr
	cmd_1.Run() //执行命令
	res_1, _ := GBK_to_UTF8(stdout.Bytes())
	res_2, _ := GBK_to_UTF8(stdout.Bytes())
	outstr, errstr := string(res_1), string(res_2)
	fmt.Println(outstr)
	fmt.Println(errstr) //错误的输出是为linux命令
}

func create_windows_process() {
	//cmd:=exec.Command("nodepad")//返回的是一个命令
	//cmd:=exec.Command("nodepad","文件的了路径")//用nodepad打开后面路径的文件
	//err:=cmd.Run()//执行命令
	cmd := exec.Command("tasklist")  //tasklist是进程列表where
	out, err := cmd.CombinedOutput() //获取命令的输出

	if err != nil {
		fmt.Println(err)
	} else {
		res, err := GBK_to_UTF8(out)
		if err != nil {
			fmt.Printf("转码失败", err)
		} else {
			fmt.Println(string(res)) //因为go是utf-8 后台输出中文是gbk所以需要转码
		}

	}
}

func process_message() {
	os.Setenv("NAME", "环境变量的值")                        //设置环境变量
	cmd := exec.Command("echo", os.ExpandEnv("$NAME")) //抓取环境变量
	cmd.Run()
	fmt.Println(os.Environ())    //获取系统的环境变量
	fmt.Println(cmd.Args)        //命令输入
	fmt.Println(cmd.Path)        //路径
	fmt.Println(cmd.Process.Pid) //进程编号
	cmd.Process.Kill()           //杀进程  等等进程信息
}

func pipe_mathod() {
	cmd := exec.Command("echo", "fdafdsafdsf")
	stdout, _ := cmd.StdoutPipe() //创建管道
	cmd.Start()

	r := bufio.NewReader(stdout)     //读取管道
	line := make([]byte, 4096, 4096) //开辟内存
	n, _ := r.Read(line)             //读取命令
	fmt.Println(string(line[:n]))    //获取进程的输出

	cmd_1 := exec.Command("echo", string(line[:n])) //一个进程的输出当做另一个进程的输入
	cmd_1.Stdout = os.Stdout
	cmd_1.Run()
}

//gbk转utf8
func GBK_to_UTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//utf8转gbk
func UTF8_to_GBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
