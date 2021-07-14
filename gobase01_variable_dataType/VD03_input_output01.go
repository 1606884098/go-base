package main

import "fmt"

/*
格式 含义
%% 一个%字面量
%b 一个二进制整数值(基数为2)，或者是一个(高级的)用科学计数法表示的指数为2的浮点数
%c 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符
%d 一个十进制数值(基数为10)
%f 以标准记数法表示的浮点数或者复数值
%o 一个以八进制表示的数字(基数为8)
%p 以十六进制(基数为16)表示的一个值的地址，前缀为0x,字母使用小写的a-f表示
%q 使用Go语法以及必须时使用转义，以双引号括起来的字符串或者字节切片[]byte，或者是以单引号括起来的数字
%s 字符串。输出字符串中的字符直至字符串中的空字符（字符串以'\0‘结尾，这个'\0'即空字符）
%t 以true或者false输出的布尔值
%T 使用Go语法输出的值的类型
%x 以十六进制表示的整型值(基数为十六)，数字a-f使用小写表示
%X 以十六进制表示的整型值(基数为十六)，数字A-F使用小写表
*/
func main() {
	var test string = "123456"
	f, _ := fmt.Scanf(test)
	fmt.Print(f)

}
