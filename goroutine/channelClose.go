package main

import (
	"fmt"
	"time"
)

// https://time.geekbang.org/column/article/14660
// 试图关闭一个已经关闭的channel会引发panic
// 接收操作可以感知到通道已经关闭，并能安全退出。
func main() {
	ch := make(chan int, 10)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch <- i
		}
		close(ch) // 关闭通道。一般使用发送方关闭通道
		fmt.Println("Sender: close the channel...")
	}()

	// 接收方。
	for {
		elem, ok := <-ch
		time.Sleep(time.Second * 1)

		// 如果ok为false，说明"通道已经关闭，并且通道内没有元素了"
		// 通道关闭后，如果还有元素未被取出，ok仍然是 true
		// 因此，通过ok判断当前通知是否关闭时间可能存在延迟
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
