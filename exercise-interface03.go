package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("已连接:", pc.name)
}

func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("未连接:", pc.name)
		return
	}
	fmt.Println("未知设备")
}

func main() {
	var a USB
	a = PhoneConnecter{"手机连接器"}
	a.Connect()
	Disconnect(a)
}
