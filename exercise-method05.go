package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increase(num TZ) {
	*tz += num
}

func main() {
	var a TZ
	a.Increase(500)
	fmt.Println(a)
}
