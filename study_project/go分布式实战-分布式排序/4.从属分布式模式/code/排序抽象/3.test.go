package main

import "math/rand"
import "strconv"
import (
	"fmt"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
AAA:
	filenamenumber := rand.Int() % 10000

	filepath := "C:\\Users\\Tsinghua-yincheng\\Desktop\\day4\\tmpfile\\"
	filestorepath := filepath + strconv.Itoa(filenamenumber) + ".txt"
	_, err := os.Stat(filestorepath)
	if err == nil {
		goto AAA
	}
	fmt.Println(filestorepath)
}
