package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println("test")
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 20; i++{
		fmt.Println(RandStringRunes(10))
	}
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}