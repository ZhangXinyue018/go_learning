package general

import (
	"fmt"
	"net"
)

func GeneralLearning(name string) {
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}
