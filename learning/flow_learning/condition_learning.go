package main

import "fmt"

func main()  {
	intA := 1
	switch {
	case intA >= 0:
		//intA ++
		fmt.Println("a>0")
	case intA == 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
	fmt.Println(intA)

LABEL1:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL1
		}
	}
}