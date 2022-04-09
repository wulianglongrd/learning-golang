package main

import "fmt"

func main() {
	mapLoop()
}

func mapLoop() {
	m := map[string]string{"foo": "11", "bar": "22"}

	for item := range m {
		fmt.Println(item)
	}

	for k, v := range m {
		fmt.Printf("%s, %s\n", k, v)
	}
}
