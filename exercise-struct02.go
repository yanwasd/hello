package main

import (
	"fmt"
)

type human struct {
	sex int
}

type teacher struct {
	Name string
	Age  int
	human
}

type student struct {
	Name string
	Age  int
	human
}

func main() {
	a := teacher{Name: "joe", Age: 30, human: human{sex: 0}}
	b := student{Name: "joe", Age: 20, human: human{sex: 1}}
	a.Name = "joe1"
	a.Age = 18
	a.human.sex = 1
	b.Name = "joe2"
	b.Age = 19
	b.sex = 0
	fmt.Println(a, b)
}
