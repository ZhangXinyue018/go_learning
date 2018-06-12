package main

import (
	"fmt"
	"time"
)

func main() {
	//StupidTest()
	//ChannelTest()
	ForRangeChannelTest()
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
