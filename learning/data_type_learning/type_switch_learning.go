package main

import (
	"fmt"
)

type PhoneConnector2 struct {
	PhoneConnector3
}

type PhoneConnector3 struct {
	name string
}

func main() {
	a := PhoneConnector2{PhoneConnector3: PhoneConnector3{"PhoneConnector2 Tester"},}
	Disconnect(a)

	b := PhoneConnector3{"PhoneConnector3 Tester"}
	Disconnect(b)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector3:
		fmt.Println("Disconnected: ", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
