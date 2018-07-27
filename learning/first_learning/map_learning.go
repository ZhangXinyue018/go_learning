package main

import (
	"fmt"
	"sort"
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


	SortTest()
}


func SortTest()(){
	a := map[string]interface{}{}
	a["a_test"] = 1
	a["m_test"] = "hello"
	a["b_test"] = 12.34

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", a[key])
	}
}
