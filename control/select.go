package main

import (
	"fmt"
	"math/rand"
	"time"
)

// more select example @see goroutine/channelSelect.go
func main() {
	abort := make(chan struct{})

	go func() {
		for {
			time.Sleep(time.Second)
			i := rand.Intn(10)
			fmt.Println("rand num =", i)
			if i > 9 {
				abort <- struct{}{}
			}
		}
	}()

	// select 一直等待，直接一次通信告知有一些情况可以执行。
	// 如果有多个情况同时发生（多个case同时满足），select会随机选择一个，这样保证每个通道有相同的机会被选中。
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("timeout")
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}

	fmt.Println("launch !")
}
