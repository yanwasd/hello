package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("i =", i)
			fmt.Println("j =", j)
			goto LABEL
		}
	LABEL:
		fmt.Println("i =", i, "\n")
	}
}
