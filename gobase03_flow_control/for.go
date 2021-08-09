package main

import "fmt"

func main() {
	var str = "woai北京"
	forFunc(str)
	forRangeFunc(str)
}

/*对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节*/
func forFunc(s string) {
	fmt.Println("普通for----------")
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i])) //出现乱码
	}
}

/*for-range 遍历方式而言，是按照字符方式遍历。因此如果有字符串有中文，也是 ok*/
func forRangeFunc(s string) {
	fmt.Println("for-range----------")
	for index, val := range s {
		fmt.Println("index=%d,val=%c", index, string(val))
	}
}
