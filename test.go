// 用iota输出计算机存储单位的例子
package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("1Byte =", B)
	fmt.Println("1KB   =", KB)
	fmt.Println("1MB   =", MB)
	fmt.Println("1GB   =", GB)
	fmt.Println("1TB   =", TB)
	fmt.Println("1PB   =", PB)
	fmt.Println("1EB   =", EB)
	fmt.Println("1ZB   =", ZB)
	fmt.Println("1YB   =", YB)
}
