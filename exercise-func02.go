package main

import (
	"fmt"
)

func main() {
	a := 1
	A(&a)
	fmt.Println("main打印出来a的值是", a)
	a = 3
	fmt.Println("main打印出来a的值是", a)
}

func A(a *int) {
	*a = 2
	fmt.Println("a的地址是：", a)
	fmt.Println("a的值是：", *a)
}
