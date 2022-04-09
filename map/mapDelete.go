package main

import "fmt"

func main() {
	mapDelete()
}

func mapDelete() {
	m2 := map[string]string{"foo": "11", "bar": "22"}
	fmt.Println(m2) // map[bar:22 foo:11]

	delete(m2, "foo")
	fmt.Println(m2) // map[bar:22]
}
