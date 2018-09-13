package main

import (
	"fmt"
	"errors"
)

type Account struct {
	Id   int
	Name string
}

func (account *Account) hasName() bool {
	fmt.Println(account.Name)
	return len(account.Name) > 0
}

func (account Account) hasName2() bool {
	fmt.Println(account.Name)
	return len(account.Name) > 0
}

func main() {
	TestStruct()
}

func General() () {
	account1 := Account{Id: 1, Name: "test",}
	fmt.Println(account1.hasName())
	fmt.Println(account1.hasName2())

	account2 := &Account{Id: 2,}
	fmt.Println(account2.hasName())
	fmt.Println(account2.hasName2())
}

type DefinedError error

type FirstError struct {
	DefinedError
}

func TestStruct() () {
	err := FirstError{DefinedError: errors.New("test lalalal")}
	err2 := err.DefinedError
	err3 := errors.New("hello")

	typeof(err)
	typeof(err2)
	typeof(err3)
}

func typeof(v interface{}) () {
	switch v.(type) {
	case FirstError:
		fmt.Println(v.(FirstError).Error())
		fmt.Println("First Error")
	case DefinedError:
		fmt.Println("Defined Error")
	default:
		fmt.Println("Not found")
	}
}
