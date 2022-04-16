package main

import "fmt"

func main()  {
 foo := Bird{Animal{"foo"}, true}
 bar := Fish{Animal{"bar"}, true}
 // 就近原则
 foo.show() // Animal: name is foo
 bar.show() // Fish: name is bar
}

type Animal struct {
  name string
}

type Bird struct {
  Animal
  canFly bool
}

type Fish struct {
  Animal
  livingSea bool
}

func (a Animal) show()  {
  fmt.Printf("Animal: name is %s\n", a.name)
}

func (f Fish) show()  {
  fmt.Printf("Fish: name is %s\n", f.name)
}
