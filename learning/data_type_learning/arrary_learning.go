package main

import "fmt"

func main()  {
	var arrayA [2]int
	arrayB := [2]int{1, 2}
	arrayC := [20]int{19: 1}
	arrayD := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arrayA)
	fmt.Println(arrayB)
	fmt.Println(arrayC)
	fmt.Println(arrayD)

	arrayE := [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}
	fmt.Println(arrayE)
}
