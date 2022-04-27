package main

import (
	"fmt"
	"time"
)

func main() {
	// 缓冲通道，缓冲队列长度为3
	ch := make(chan int, 3)

	no := 0
	go func() {
		for {
			time.Sleep(time.Second * 2)
			no++
			ch <- no
			fmt.Println("provider: ", no)
		}
	}()

	for {
		time.Sleep(time.Second * 3)
		fmt.Println("consume: ", <-ch)
	}
}
