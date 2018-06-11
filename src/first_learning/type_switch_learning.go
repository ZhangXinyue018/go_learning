package main

import (
	"fmt"
)

type PhoneConnector2 struct {
	name string
}

type PhoneConnector3 struct {
	name string
}

func main() {
	a := PhoneConnector2{"PhoneConnector2 Tester"}
	Disconnect(a)

	b := PhoneConnector3{"PhoneConnector3 Tester"}
	Disconnect(b)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector2:
		fmt.Println("Disconnected: ", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
