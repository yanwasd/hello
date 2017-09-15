package main

import (
	"fmt"
)

type human struct {
	sex  int
	Name string
}

type contact struct {
	phone int
	Name  string
}

type teacher struct {
	Age int
	human
	contact
}

type student struct {
	Name string
	Age  int
	human
}

func main() {
	a := teacher{Age: 30, human: human{sex: 0, Name: "joe3"}}
	b := student{Name: "joe", Age: 20, human: human{sex: 1}}
	//a.Name = "joe1"
	a.Age = 18
	a.human.sex = 1
	a.human.Name = "joe4"
	a.contact.phone = 13032032610
	a.contact.Name = "joe6"
	b.Name = "joe2"
	b.Age = 19
	b.sex = 0
	b.human.Name = "joe5"
	fmt.Println(a.human.Name, b)
}
