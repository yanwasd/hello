package main

import (
	"fmt"
)

func main() {
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}
}
