package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GoCmd(cmdstr string) string {
	cmd := exec.Command(cmdstr)
	cmd.Run()
	return ""
}

func GoCmdWithResult(cmdstr string) string {
	cmdstr = strings.TrimSpace(cmdstr)
	cmd := exec.Command(cmdstr)

	stdout, err := cmd.StdoutPipe() //
	if err != nil {
		return "error1"
	}
	if err := cmd.Start(); err != nil {
		return "error2"
	}
	outbyte, err := ioutil.ReadAll(stdout) //读取所有的输出
	stdout.Close()                         //关闭
	if err := cmd.Wait(); err != nil {
		return "error3"
	}
	return string(outbyte) //返回结果

}

func main() {
	taskcmd := GoCmd("tasklist")
	fmt.Println(taskcmd)
	task := GoCmdWithResult("ipconfig  ")
	fmt.Println(task)
}
