package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	philosopherMu    sync.Mutex
)

type PhilosopherGroup struct {
	Vacancy int
}

func (p *PhilosopherGroup) leftPick(name string) bool {
	philosopherMu.Lock()
	defer philosopherMu.Unlock()
	if p.Vacancy > 0 {
		p.Vacancy--
		return true
	} else {
		return false
	}
}

func (p *PhilosopherGroup) eat(name string) {
	fmt.Printf("%s is eating with left vacancy %d\n", name, p.Vacancy)
}

func (p *PhilosopherGroup) rightLeave(name string) {
	philosopherMu.Lock()
	defer philosopherMu.Unlock()
	p.Vacancy++
}

func (p *PhilosopherGroup) TryEat(name string) {
	for !p.leftPick(name){
		continue
	}
	p.eat(name)
	time.Sleep(5 * time.Second)
	p.rightLeave(name)
}

func main() {
	pg := &PhilosopherGroup{5}
	for i := 0; i < 10; i++ {
		go pg.TryEat(strconv.Itoa(i))
	}
	time.Sleep(60 * time.Second)
}