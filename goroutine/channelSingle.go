package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//channelType()
	singleChan()
}

// 通道类型
func channelType() {
	var ch = make(chan int)         // 双向通道
	var producer = make(chan<- int) // 只能发，不能收
	var consumer = make(<-chan int) // 只能收，不能发

	fmt.Printf("%T\n", ch)       // chan int
	fmt.Printf("%T\n", producer) // chan<- int
	fmt.Printf("%T\n", consumer) // <-chan int
}

// 单向通道最主要的用途就是约束其他代码的行为
func singleChan() {
	ch := make(chan int)
	go chanProducer(ch)
	go chanConsumer(ch)

	select {} // 主线程永远阻塞
}

func chanProducer(ch chan<- int) {
	for {
		no := rand.Intn(10)
		ch <- no
		fmt.Println("producer: ", no)
		time.Sleep(time.Second)
	}
}

func chanConsumer(ch <-chan int) {
	for {
		fmt.Println("consumer: ", <-ch)
	}
}
