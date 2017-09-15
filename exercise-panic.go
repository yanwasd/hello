package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("This is func A")
}
func B() {
	panic("This is panic in func B")
}
func C() {
	fmt.Println("This is func C")
}
