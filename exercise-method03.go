package main

import (
	"fmt"
)

type A struct {
	name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.name)
}

func (a *A) Print() {
	a.name = "yanwasd"
	fmt.Println(a.name)
}
