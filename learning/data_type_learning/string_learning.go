package main

import (
	"fmt"
	"strconv"
)

type TestA struct {
	Abc int
}

func main() {
	Test2()
}

func Test1() () {
	var a string
	fmt.Println(len(a))

	var testA TestA
	fmt.Printf("%p\n", &testA)

	var stringA = "1"
	intA, _ := strconv.Atoi(stringA)
	fmt.Println(intA)
}

func Test2() () {
	// 在golang中字符串的每一个字符所占的字节数并不是一定的
	test := "这是一个aaaaa的奇怪测试"
	for index, value := range test{
		fmt.Printf("%d: %s \n", index, string(value))
	}
	fmt.Println(len(test))
}
