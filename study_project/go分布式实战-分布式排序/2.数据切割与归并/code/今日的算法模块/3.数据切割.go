package main

import "os"
import "io"
import (
	"bufio"
	"fmt"
	"strconv"
)

func GetLineNumber(filepath string) int {
	fi, err := os.Open(filepath)
	defer fi.Close() //打开文件
	if err != nil {
		return -1
	}
	i := 0 //统计行数
	br := bufio.NewReader(fi)
	for {
		_, _, err := br.ReadLine()
		if err == io.EOF {
			break //文件结束，末尾
		}
		i++ //统计行数
	}
	return i

}
func evgSplit(num, N int) []int {
	arr := []int{}
	if num%N == 0 {
		for i := 0; i < N; i++ {
			arr = append(arr, num/N)
		}
	} else {
		evg := (num - num%N) / (N - 1) //97 -7 //9=10
		for i := 0; i < N-1; i++ {
			arr = append(arr, evg) //追加
			num -= evg
		}
		arr = append(arr, num) //7

	}
	return arr
}

func main() {
	filepath := "Z:\\E\\lastvedio5\\day1综合project练习\\chinaunix_com.txt"
	N := GetLineNumber(filepath)
	fmt.Println(N)
	var num = 7
	arrlist := evgSplit(N, num) //数据切割为7份

	fi, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi) //读取

	filesavepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day2\\data\\"
	for i := 0; i < len(arrlist); i++ {
		fmt.Println(i)
		tmppath := filesavepath + "chinaunix" + strconv.Itoa(i) + ".txt"
		savefile, _ := os.Create(tmppath)
		defer savefile.Close()
		save := bufio.NewWriter(savefile)
		for j := 0; j < arrlist[i]; j++ {
			line, _, _ := br.ReadLine()
			fmt.Fprintln(save, string(line)) //写入
		}
		save.Flush() //刷新
	}

}
