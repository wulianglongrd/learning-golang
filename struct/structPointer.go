package main

import "fmt"

func main() {
	structPointer()
}

type Book struct {
	title  string
	bookId int
}

func structPointer() {
	b := Book{title: "go in action", bookId: 1}
	var p *Book
	p = &b
	fmt.Println(p)  // &{go in action 1}
	fmt.Println(*p) // {go in action 1}

	fmt.Println(p.title, p.bookId) // go in action 1

	fmt.Println()

	p.title = "english book"
	fmt.Println(p) // &{english book 1}
}
