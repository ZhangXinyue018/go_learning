package main

import (
	"fmt"
	"sync"
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
}

func ChangeTemp(temp *Temp){
	fmt.Println(temp)
	temp.Test = 5
	fmt.Println(temp)
	wg.Done()
}