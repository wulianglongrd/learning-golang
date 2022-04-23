package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			q.Dequeue()
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 3)
}

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
	fmt.Println("put item to queue")

	q.cond.Broadcast()
}

func (q Queue) Dequeue() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		fmt.Println("no data, wait")
		q.cond.Wait()
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	fmt.Println("get value: ", item)
}
