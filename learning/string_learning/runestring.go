package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	TestRune()
}

func TestRune() {
	s := "这是一个aaa测试的字符串!"
	fmt.Println(utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		result, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", result)
		//fmt.Println()
	}
	fmt.Println("================================")
	fmt.Println(len([]rune(s)))
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) \n", i, ch)
	}

	fmt.Println("================================")
	fmt.Println(len([]byte(s)))
	for i, ch := range []byte(s) {
		fmt.Printf("(%d, %c) \n", i, ch)
	}

	fmt.Println("================================")
	fmt.Println(len(s))
	for i, str := range s {
		fmt.Printf("%d: %s \n", i, string(str))
	}

}
