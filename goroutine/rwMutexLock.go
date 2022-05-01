package main

import (
	"fmt"
	"sync"
	"time"
)

// https://golang.google.cn/pkg/sync/#RWMutex
// https://www.qfgolang.com/?special=bingfagoroutinechannel&pid=2132
// 读写互斥锁
// 1. 同时只能有一个 goroutine 能够获得写锁
// 2. 同时可以有任意多个 goroutine 获得读锁
// 3. 同时只能存在写锁 或 读锁（读和写锁互斥）
// 常用于读次数远多于写次数的场景

var z = 0

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}

	wg.Add(3)

	go writeData(&mutex, &wg)
	go readData(&mutex, &wg)
	go writeData(&mutex, &wg)

	wg.Wait()
}

func readData(mutex *sync.RWMutex, wg *sync.WaitGroup) {
	mutex.RLock()
	defer mutex.RUnlock()

	fmt.Println("read z =", z)
	time.Sleep(time.Second * 3)
	wg.Done()
}

func writeData(mutex *sync.RWMutex, wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()

	z = z + 1
	fmt.Println("write z =", z)
	time.Sleep(time.Second * 3)
	wg.Done()
}
