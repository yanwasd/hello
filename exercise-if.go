package main

import (
	"fmt"
)

func main() {
	a := true
	if a, b, c := 1, 2, 3; a+b+c > 6 {
		fmt.Println("a,b,c之和大于6")
	} else {
		fmt.Println("a,b,c之和小于等于6")
		fmt.Println("里面的a值是：", a)
	}
	fmt.Println("外面的a值是：", a)
}
