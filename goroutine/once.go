package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := sync.Once{}
	go once.Do(runOnce)
	go once.Do(runOnce)
	go once.Do(runOnce)

	time.Sleep(time.Second)
}

func runOnce() {
	fmt.Println("run...")
}
