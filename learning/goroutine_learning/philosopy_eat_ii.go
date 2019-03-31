package main

import (
	"golang.org/x/exp/errors/fmt"
	"strconv"
	"time"
)

type Philosopy struct {
}

func (p *Philosopy) TryEat(name string, left chan int) {
	for {
		select {
		case <-left:
			fmt.Printf("%s is eating\n", name)
			time.Sleep(2 * time.Second)
			left <- 1
			return
		}
	}
}

func main() {
	left := make(chan int, 10)
	go func() {
		for i := 0; i < 5; i++ {
			left <- 1
		}
	}()
	for i := 0; i < 10; i++ {
		p := &Philosopy{}
		go p.TryEat(strconv.Itoa(i), left)
	}

	time.Sleep(60 * time.Second)
}
