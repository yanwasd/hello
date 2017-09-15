package main

import (
	"fmt"
)

type A int

func main() {
	var a A
	a = 1
	a.Increase()
}

func (a A) Increase() {
	a = a + 100
	fmt.Println(a)
}
