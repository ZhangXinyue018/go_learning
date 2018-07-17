package main

import "fmt"

type Account struct{
	Id int
	Name string
}


func (account *Account) hasName() bool{
	fmt.Println(account.Name)
	return len(account.Name) > 0
}

func (account Account) hasName2() bool{
	fmt.Println(account.Name)
	return len(account.Name) > 0
}

func main()  {
	account1 := Account{Id:1, Name:"test",}
	fmt.Println(account1.hasName())
	fmt.Println(account1.hasName2())

	account2 := &Account{Id:2, }
	fmt.Println(account2.hasName())
	fmt.Println(account2.hasName2())
}
