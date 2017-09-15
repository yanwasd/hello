package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	A(a, b)
	fmt.Println("main打印出来a,b的值是", a, b)
}

func A(s ...int) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}
