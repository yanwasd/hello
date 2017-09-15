package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("for循环打印的i值为：", i)
		defer func() {
			fmt.Println("defer后闭包函数打印的i值为：", i)
		}()
	}
}
