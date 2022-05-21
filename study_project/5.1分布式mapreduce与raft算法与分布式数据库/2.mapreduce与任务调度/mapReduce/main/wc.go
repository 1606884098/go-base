package main

import (
	"eth-1804/mapReduce/c04-distribute"
)

func MapFunc(file string, value string) (res []c04_distribute.KeyValue) {

	return
}

// 自定义的reduce聚合函数
func ReduceFunc(key string, values []string) string {

	return ""
}
