package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	PI        float32 = 3.14
	testConst         = "2"
	testA             = 1
	testB
	testC
)

const (
	testD, testE = 1, "2"
	testF        = 3
)

const (
	Monday    = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Hello Go!中文测试")
	fmt.Println(PI)
	fmt.Println(testConst)
	fmt.Println(math.MaxUint8)
	var a float32 = 65.9
	b := int(a)
	fmt.Println(b)
	fmt.Println(strconv.Itoa(b))
	fmt.Println(string(b))
	fmt.Println(testA)
	fmt.Println(testB)
	fmt.Println(testC)
	fmt.Println(testD)
	fmt.Println(testE)
	fmt.Println(testF)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)
}
