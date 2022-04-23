package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitBySleep()
	//waitByChannel()
	waitByWaitGroup()
}

// 使用sleep防止主线程退出
func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	fmt.Println("waitBySleep done")
	time.Sleep(time.Second)
}

func waitByChannel() {
	ch := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- i
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-ch
	}
	fmt.Println("waitByChannel done")
}

func waitByWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("waitByWaitGroup done")
}
