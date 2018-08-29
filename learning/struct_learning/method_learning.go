package main

import "fmt"

func main() {
	a := TestMethodType{Name: "test", Age: 19,}
	a.Print()
	(*TestMethodType).Print(&a)
	fmt.Println(a)
	a.Add()
	fmt.Println(a)
}

type TestMethodType struct {
	Name string
	Age  int
}

func (test TestMethodType) Print() {
	fmt.Println("test method")
}

func(test *TestMethodType) Add(){
	test.Age ++
}
