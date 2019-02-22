package main

import (
	"fmt"
)

type TestA struct {
	Abc int
}

func main() {
	Test1()
}

func Test1() () {
	var a []int
	fmt.Println(len(a))
}

func Test2() () {
	// 在golang中字符串的每一个字符所占的字节数并不是一定的
	test := "这是一个aaaaa的奇怪测试"
	for index, value := range test{
		fmt.Printf("%d: %s \n", index, string(value))
		//fmt.Println(reflect.TypeOf(value))
	}
	fmt.Println(len(test))
}
