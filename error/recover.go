package main

import "fmt"

// https://go.dev/blog/defer-panic-and-recover
// https://gobyexample.com/recover
func main() {
	controller()
}

func controller() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered, err:", r)
		}
	}()

	panic("a problem")
	fmt.Println("after panic")
}
