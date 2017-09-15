package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}

func A(per person) {
	per.Age = 13
	fmt.Println("A", per)
}
