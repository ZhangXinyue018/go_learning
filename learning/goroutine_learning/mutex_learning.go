package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	wg    sync.WaitGroup
	value = 0
)

func Test(i int) {
	defer wg.Done()
	mu.Lock()
	defer func() {
		mu.Unlock()
		//fmt.Printf("unlocked %d \n", i)
	}()
	fmt.Printf("locking of %d\n", i)
	value++
	fmt.Printf("value is %v \n", value)
}

func main() {
	count := 100
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go Test(i)
	}
	wg.Wait()
}
