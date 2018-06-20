package main

import "fmt"

type TestA struct{
	Abc int
}

func main()  {
	var a string
	fmt.Println(len(a))

	var testA TestA
	fmt.Printf("%p", &testA)
}
