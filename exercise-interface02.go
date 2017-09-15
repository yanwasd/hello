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
	fmt.Println("Connected:", pc.name)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnected.")
}

func main() {
	var a USB
	a = PhoneConnecter{"手机连接器"}
	a.Connect()
	Disconnect(a)
}
