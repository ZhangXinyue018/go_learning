package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("Test")
	}
	a()
	b := closure(1)
	fmt.Println(b(2))
	fmt.Println(b(3))

}

func FirstFunc(a int, b string) (int, string, int) {
	return a, b, 1
}

func SecondFunc(a int, b string) int {
	return 1
}

func ThirdFunc() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func FourthFunc(a ... int) {
	return
}

func closure(x int) (func(int) int) {
	return func(y int) int {
		return x + y
	}
}
