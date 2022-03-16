package main

import "regexp"
import "fmt"

func getnumberstr(str1 string) string {
	//str1:="163-1.txt"
	reg := regexp.MustCompile(`(\d+)`)
	ss := reg.FindAllStringSubmatch(str1, -1)
	//fmt.Println(strconv.Itoa(int(ss[0][1]-48)))
	fmt.Println("ss---", ss, ss[0][1])
	return ss[0][1]
}
func main() {
	fmt.Println(getnumberstr("C:\\Users\\Tsinghua-yincheng\\Desktop\\day5\\tmp\\csdnPasswd23sort.txt"))
}
