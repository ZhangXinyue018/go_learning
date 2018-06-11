package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type contact struct {
	Phone, Email string
}

type fullPerson struct {
	Name string
	Age  int
	contact
}

func main() {
	a := &person{
		Name: "test",
		Age:  16,
	}
	a.Name = "changed name"
	fmt.Println(a)
	ChangeAge(a)
	fmt.Println(a)

	b := &struct {
		Name string
		Age  int
	}{
		Name: "second",
		Age:  19,
	}
	fmt.Println(b)

	c := &fullPerson{
		Name:    "test",
		Age:     16,
		contact: contact{Email: "aaa@aaa.com"},
	}
	c.Email = "email@aaa.com"
	c.Phone = "1234567"
	fmt.Println(*c)

	d := person{
		Name: "testName",
		Age:  16,
	}

	e := d
	e.Name = "testName1234"
	fmt.Println(d)
	fmt.Println(e)
}

func ChangeAge(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}
