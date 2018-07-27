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

func Test1() () {
	newTest := GetTest()
	defer newTest.Hello()
	name := newTest.Name
	fmt.Println("get name" + name)
}

func Test2() () {
	defer func() {
		if x := recover(); x != nil {
			//panic("test")
			fmt.Println("ignore panic")
		}
	}()
	panic("ENTER")
}

func main() {
	Test2()
}
