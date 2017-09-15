package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("这是func A")
	}
	a()
}
