package main

import (
	"fmt"
	"strconv"
)

type TestA struct {
	Abc int
}

func main() {
	var a string
	fmt.Println(len(a))

	var testA TestA
	fmt.Printf("%p\n", &testA)

	var stringA = "1"
	intA, _ := strconv.Atoi(stringA)
	fmt.Println(intA)
}
