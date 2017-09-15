package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{0: "j", 8: "h", 3: "c", 9: "i", 5: "e", 7: "g", 6: "f", 4: "d", 2: "b", 1: "a"}
	fmt.Println("m1数组的值为：", m1)
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println("m2数组的值为：", m2)
}
