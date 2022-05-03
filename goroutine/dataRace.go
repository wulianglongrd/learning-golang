package main

import (
	"fmt"
	"time"
)

// $ go run -race dataRace.go
// main goroutine... 3
// ==================
// WARNING: DATA RACE
// Write at 0x00c00001a0f8 by goroutine 7:
// main.main.func1()
// /Users/logx/go/src/github.com/wulianglongrd/learning-golang/goroutine/dataRace.go:31 +0x34
//
// Previous write at 0x00c00001a0f8 by main goroutine:
// main.main()
// /Users/logx/go/src/github.com/wulianglongrd/learning-golang/goroutine/dataRace.go:35 +0xb8
//
// Goroutine 7 (running) created at:
// main.main()
// /Users/logx/go/src/github.com/wulianglongrd/learning-golang/goroutine/dataRace.go:30 +0xae
// ==================
// goroutine... 2
// Found 1 data race(s)
// exit status 66

func main() {
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine...", a)
	}()

	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine...", a)
}
