package main

import (
	"fmt"
	"github.com/yanwasd/stringutil"
	"strconv"
)

func main() {
	fmt.Printf("你好，世界！\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("\n")
	var a int = 67
	fmt.Println("a =", a)
	b := string(a)
	fmt.Println("b = ", b)
	c := strconv.Itoa(a)
	fmt.Println("c = ", c)
	d, _ := strconv.Atoi(c)
	fmt.Println("d = ", d)
}
