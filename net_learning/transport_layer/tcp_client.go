package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	temp := make(chan int)
	for i := 1; i < 10; i++ {
		go TcpDial()
	}
	<- temp
}

func TcpDial() {
	out, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	//bs, _ := ioutil.ReadAll(out)
	//fmt.Println(string(bs))
	io.WriteString(out, fmt.Sprintf("start %v \n", time.Now()))
	time.Sleep(5*time.Second)
	io.WriteString(out, fmt.Sprintf("end %v \n", time.Now()))
}
