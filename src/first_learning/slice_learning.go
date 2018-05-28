package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a
	s1 := a[5:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)

	a[5], a[6], a[7], a[8] = 0, 0, 0, 0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)

	s1[3] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s1)


	s2 := make([]int, 3, 6)
	fmt.Println(len(s2), cap(s2))
	fmt.Printf("%v %p\n", s2, s2)
	s2 = append(s2, 2, 4)
	fmt.Println(len(s2), cap(s2))
	fmt.Printf("%v %p\n", s2, s2)
	s2 = append(s2, 22, 24)
	fmt.Println(len(s2), cap(s2))
	fmt.Printf("%v %p\n", s2, s2)
}
