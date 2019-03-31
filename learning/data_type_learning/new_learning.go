package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	TempValue = new(Temp)
	wg = sync.WaitGroup{}
)

type Temp struct {
	Test int
}

func main() {
	wg.Add(2)
	go ChangeTemp(TempValue)
	go ChangeTemp(TempValue)
	wg.Wait()
	//MockContextCase()
}

func ChangeTemp(temp *Temp){
	fmt.Println(temp)
	temp.Test = 5
	fmt.Println(temp)
	wg.Done()
}


func MockContextCase(){
	fmt.Println(time.Now())
	NewTestOutside()
	fmt.Println(time.Now())
}

func NewTestOutside(){
	go NewTest()
	time.Sleep(5*time.Second)
}

func NewTest(){
	time.Sleep(10 * time.Second)
	fmt.Println("I'm done")
}