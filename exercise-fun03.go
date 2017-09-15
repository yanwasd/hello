package main

import (
	"fmt"
)

func main() {
	a := A
	a()
}

func A() {
	fmt.Println("这是func A")
}
