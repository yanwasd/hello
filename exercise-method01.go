package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()

	b := B{}
	b.Print()
}

func (a A) Print() {
	fmt.Println("A")
}

func (b B) Print() {
	fmt.Println("B")
}
