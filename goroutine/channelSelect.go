package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//selectBasic()
	//selectTimeout()
	//selectDefault()
	selectSignalOnly()
}

// Break out of select loop?
// https://stackoverflow.com/questions/25469682/break-out-of-select-loop
func selectSignalOnly() {
	ch := make(chan string)
	done := make(chan bool)

	go time.AfterFunc(time.Second*5, func() {
		done <- true
	})

	go func() {
		for {
			time.Sleep(time.Second)
			ch <- "bomb news!!!"
		}
	}()

doneLabel:
	for {
		select {
		case news := <-ch:
			fmt.Println("receive news:", news)
		case <-done:
			fmt.Println("get stop single")
			break doneLabel
		}
	}

	fmt.Println("done")
}

// https://golangbot.com/select/
func selectDefault() {
	ch := make(chan string)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			ch <- "bomb news!!!"
		}
	}()

	for {
		time.Sleep(time.Second)
		select {
		case v := <-ch:
			fmt.Println("receive news:", v)
		default:
			fmt.Println("no news")
		}
	}
}

// https://yourbasic.org/golang/select-explained/
func selectTimeout() {
	ch := make(chan string)
	secondCount := 0

	go func() {
		for {
			time.Sleep(time.Second)
			secondCount++
		}
	}()

	go func() {
		for {
			i := time.Duration(rand.Intn(10))
			fmt.Printf("second: %d, will wait seconds %d\n", secondCount, i)
			time.Sleep(time.Second * i)

			fmt.Printf("second: %d, send news\n", secondCount)
			ch <- "bomb news!!!"
		}
	}()

	for {
		select {
		case news := <-ch:
			fmt.Printf("second: %d, receive news: %s\n", secondCount, news)
		case <-time.After(time.Second * 3):
			fmt.Printf("second: %d, time out\n", secondCount)
		}
	}

}

// https://gobyexample.com/select
func selectBasic() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "tow"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}
}
