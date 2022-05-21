package main

import (
	"eth-1804/mapReduce/c01-single-basic"
	"fmt"
)

func main() {
	fmt.Println(c01_single_basic.IHASH("10") % 3)
}
