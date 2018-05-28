package main

import (
	"fmt"
)

func main()  {
	m := make(map[int]string)
	fmt.Println(m)
	m[1] = "ok"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)
	a := m[3]
	fmt.Println(a)


	m1 := make(map[int]map[int]string)
	m1[1] = make(map[int]string)
	m1[1][1] = "ok"
	fmt.Println(m1)
	//b, ok := m1[2][1]
	//fmt.Println(b, ok)
}
