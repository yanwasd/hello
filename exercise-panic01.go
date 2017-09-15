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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("This is recover in func B")
		}
	}()
	panic("This is panic in func B")
}
func C() {
	fmt.Println("This is func C")
}
