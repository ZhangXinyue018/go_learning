package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
	Test
}

type Test struct {
	TestName string
}

func (u User) Hello() {
	fmt.Println("Hello World!")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("XX")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v ", f.Name, f.Type, val)
		fmt.Printf("%#v\n", f)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v \n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "OK", 12, Test{"HAHAHAHHA"}}
	Info(u)
	
}
