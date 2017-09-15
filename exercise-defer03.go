package main

import (
	"fmt"
)

func main() {
	whatever := []int{5, 6, 7, 8}

	fmt.Println("数组的内容是：", whatever)

	for i, j := range whatever {
		fmt.Printf("第1个for循环打印的i值是：%d,j值是： %d \n", i, j)
	}

	for i, j := range whatever {
		fmt.Printf("第2个for循环打印的i值是：%d,j值是： %d \n", i, j)
		defer func() {
			fmt.Printf("第2个for循环里defer后的空参数匿名函数打印的i值为：%d,j值为：%d \n", i, j)
		}()
	}

	for i, j := range whatever {
		defer func(m, n int) {
			fmt.Printf("第3个for循环里defer后以i和j为参数的匿名函数打印的m值：%d,n值：%d\n", m, n)
		}(i, j)
	}
}
