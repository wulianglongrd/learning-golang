package main

import (
	"sync"
	"time"
)

func main() {
	//unsafeWrite()
	safeWrite()

	time.Sleep(time.Second)
}

// fatal error: concurrent map writes
func unsafeWrite() {
	m := map[string]int{}
	for i := 0; i < 1000; i++ {
		go func() {
			m["foo"] = i
		}()
	}
}

func safeWrite() {
	m := SafeMap{
		mp:    map[string]int{},
		Mutex: sync.Mutex{},
	}
	for i := 0; i < 1000; i++ {
		go func() {
			m.Write("foo", i)
		}()
	}
}

type SafeMap struct {
	mp map[string]int
	sync.Mutex
}

func (m *SafeMap) Write(k string, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v
}

func (m *SafeMap) Read(k string) (int, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.mp[k]
	return v, ok
}
