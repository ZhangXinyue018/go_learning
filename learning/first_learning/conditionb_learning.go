package main

import "fmt"

func main()  {
	switch intA := 1; {
	case intA >= 0:
		fmt.Println("a>0")
		fallthrough
	case intA == 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

LABEL1:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL1
		}
	}
}