package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func SplitFile(filepath string, savepath string, num int) []string {
	filelist := strings.Split(filepath, "\\")
	filename := filelist[len(filelist)-1] //取得文件名
	fmt.Println("filename", filename)

	N := GetLineNumber(filepath) //文件的行数
	arrlist := evgSplit(N, num)  //数据切割为7份

	fi, err := os.Open(filepath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		return []string{}
	}
	br := bufio.NewReader(fi) //读取

	filesavepath := savepath
	tmpfilelist := []string{}
	for i := 0; i < len(arrlist); i++ {
		fmt.Println(i)
		tmppath := filesavepath + strings.Replace(filename, ".txt", "", -1) + strconv.Itoa(i) + ".txt"
		tmpfilelist = append(tmpfilelist, tmppath) //追加路径
		savefile, _ := os.Create(tmppath)
		defer savefile.Close()
		save := bufio.NewWriter(savefile)
		for j := 0; j < arrlist[i]; j++ {
			line, _, _ := br.ReadLine()
			fmt.Fprintln(save, string(line)) //写入
		}
		save.Flush() //刷新
	}

	fi.Close()

	return tmpfilelist
}
