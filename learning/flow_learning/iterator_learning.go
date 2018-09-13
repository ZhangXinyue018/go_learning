package main

import (
	"fmt"
	"sort"
)

func main() {

	sm := make([]map[int]string, 5)

	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	testM := map[int]string{1: "a", 2: "b", 3: "c"}
	testNewM := make(map[string]int)
	for k, v :=range testM{
		testNewM[v] = k
	}
	fmt.Println(testNewM)
}
