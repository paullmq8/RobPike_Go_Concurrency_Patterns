package main

import (
	"fmt"
)

type myInt int

func (mi *myInt) print() {
	println("aaa")
}

func main() {
	var a *myInt
	fmt.Printf("%T %p %v\n", a, a, a)
	a.print()
}
