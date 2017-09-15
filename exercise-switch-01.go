package main

import (
	"fmt"
)

func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
	case a >= 1:
		fmt.Println("a>=1")
	}
	fmt.Println("a的值是：", a)
}
