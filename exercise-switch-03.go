package main

import (
	"fmt"
)

func main() {
	a := 10
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	}
	fmt.Println("switch语句上面a的值是：", a)
}
