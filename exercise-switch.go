package main

import (
	"fmt"
)

func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)
}
