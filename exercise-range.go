package main

import (
	"fmt"
)

func main() {
	//当对slice 或者array 做循环时，range 返回序号作为键，这个序号对应的内容作为值。
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Println(k, "=", v)
	}
	//字符串上直接使用range。这样字符串被打散成独立的Unicode 字符并且起始位按照UTF-8 解析。
	for pos, char := range "我要测试Golang" { //took 2 bytes
		fmt.Println(pos, "=", char)
	}
}
