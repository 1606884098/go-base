package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	/*	d := time.Now().Add(50 * time.Millisecond)
		ctx, cancel := context.WithDeadline(context.Background(), d)

		// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
		// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
		defer cancel()

		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}*/
}
