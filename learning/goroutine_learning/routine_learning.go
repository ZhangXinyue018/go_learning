package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	//StupidTest()
	//ChannelTest()
	//ForRangeChannelTest()
	ForSelectChannelTest()
}

func StupidTest() {
	go func() {
		fmt.Println("Go Coming!!!")
	}()
	time.Sleep(2 * time.Second)
}

func ChannelTest() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go channel!!!")
		c <- true
	}()
	<-c
	close(c)
}

func ForRangeChannelTest() {
	// make的channel是双向通道
	c := make(chan bool)
	go func() {
		fmt.Println("Go for range channel!!!")
		c <- true
		// 如果使用for range，必须关闭channel并保证关闭成功，否则会形成死锁
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func ForSelectChannelTest() {
	c := make(chan bool, 10) //非阻塞
	remark := make(chan bool) //阻塞
	go func() {
		for {
			select {
			case temp := <-c:
				fmt.Println(temp)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			c <- rand.Float64() > 0.5
		}
	}()
	<-remark
}
