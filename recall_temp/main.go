package main

import "fmt"

func main() {
}

func Test1() {
	t := Test{Temp: 0}
	fmt.Println(t)
	fmt.Println("==============")
	t.Try1()
	fmt.Println("==============")
	t.Try2()
}

type Test struct {
	Temp int
}

func (t Test) Try1() {
	fmt.Println("enter try 1")
	fmt.Println(t)
	t.Temp = 2
	fmt.Println(t)
}

func (t *Test) Try2() {
	fmt.Println("enter try 2")
	fmt.Println(t)
	t.Temp = 3
	fmt.Println(t)
}
