package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("for循环打印的i值为：", i)
		defer fmt.Println("defer打印的i值为：", i)

	}
}
