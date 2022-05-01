package main

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0

// https://golangbot.com/mutex/
func main() {
	//mutexLockDemo()
	channelDemo()
}

func mutexLockDemo() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// 如果 m 是按值传递而不是传递地址，每个 Goroutine 都会有自己的互斥量副本，仍然会出现竞态条件
		// If the mutex is passed by value instead of passing the address,
		// each Goroutine will have its own copy of the mutex
		// and the race condition will still occur.
		go incrWithLock(&wg, &m)
	}

	wg.Wait()
	fmt.Println("final value of x", x)
}

func incrWithLock(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()

	x = x + 1
	wg.Done()
}

// ---------------------
func channelDemo() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		incrWithChannel(&wg, ch)
	}

	wg.Wait()
	fmt.Println("final value of y", y)
}

func incrWithChannel(wg *sync.WaitGroup, ch chan bool) {
	// the buffered channel has a capacity of 1,
	// all other Goroutines trying to write to this channel are blocked
	// until the value is read from this channel after incrementing x.
	// Effectively this allows only one Goroutine to access the critical section.
	ch <- true
	y = y + 1
	<-ch
	wg.Done()
}
