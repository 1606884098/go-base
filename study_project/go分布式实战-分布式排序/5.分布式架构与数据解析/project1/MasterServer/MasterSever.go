package main

import "fmt"
import "os"
import (
	"net"
	"strings"
	"sync"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
func main1() {
	//输入文件路径
	path1 := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\csdnPasswd.txt"
	//path2:="C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\csdnPasswdRmDuplicatemidd.txt"
	//结果文件路径
	pathresult := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\"
	iplist := []string{"127.0.0.1:8848", "127.0.0.1:8849", "127.0.0.1:8850"}
	//切割
	N := len(iplist)
	savelist := SplitFile(path1, pathresult, N)
	fmt.Println("切割完成", savelist) //切割有效的
	//建立链接池
	connlist := []net.Conn{} //连接池

	for i := 0; i < len(iplist); i++ {
		tcpaddr, err := net.ResolveTCPAddr("tcp4", iplist[i])
		CheckError(err)
		conn, err := net.DialTCP("tcp", nil, tcpaddr)
		CheckError(err)
		connlist = append(connlist, conn)

	}
	//发送并发
	for i := 0; i < len(connlist); i++ {
		go func(i int) {
			SendFile(savelist[i], connlist[i], 3, true, true)
			fmt.Println("发送完成", connlist[i])
		}(i)
	}

	//接收并发
	var wg sync.WaitGroup
	wg.Add(3)
	lastsavelist := []string{}
	for i := 0; i < len(connlist); i++ {
		savepath := strings.Replace(savelist[i], ".txt", "sort.txt", -1)
		lastsavelist = append(lastsavelist, savepath)
		go func(i int) {
			ReceFile(savepath, connlist[i], 3)
			fmt.Println("发送完成", connlist[i])
			wg.Done()
		}(i)
	}
	wg.Wait() //等待接收完成之后
	fmt.Println("接收完成", lastsavelist)

	fmt.Println(Merge(lastsavelist, pathresult, 3, true, "csdnPasswd"))
	fmt.Println("归并完成")
	// waitGroup   归并

}

func main() {

	pathresult := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\"
	lastsavelist := []string{"C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\csdnPasswd0sort.txt",
		"C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\csdnPasswd1sort.txt",
		"C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\csdnPasswd2sort.txt"}
	path1 := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\csdnPasswd.txt"
	//csdnPasswd
	pathlist := strings.Split(path1, "\\")
	reslist := strings.Split(pathlist[len(pathlist)-1], ".")

	fmt.Println(Merge(lastsavelist, pathresult, 3, false, reslist[0]))
	fmt.Println("归并完成")
}
