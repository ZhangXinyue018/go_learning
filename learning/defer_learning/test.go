package main

import "fmt"

func GetTest() (*Test) {
	test := Test{Name: "123"}
	fmt.Println("enter test")
	//defer test.Hello()
	fmt.Println("return test")
	return &test
}

type Test struct {
	Name string
}

func (test *Test) Hello() () {
	fmt.Println(test.Name)
}

func main() {
	newTest := GetTest()
	defer newTest.Hello()
	name := newTest.Name
	fmt.Println("get name" + name)
}
