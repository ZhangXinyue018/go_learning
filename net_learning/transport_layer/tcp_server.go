package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func main(){
	in, err := net.Listen("tcp", ":9000")
	if err != nil{
		panic(err)
	}
	defer in.Close()
	for {
		conn, err := in.Accept()
		if err != nil{
			panic(err)
		}
		fmt.Printf("connection accepted at %v", time.Now())
		//io.WriteString(conn, fmt.Sprintf("Hello world %v \n", time.Now()))
		bs, _ := ioutil.ReadAll(conn)
		fmt.Println(string(bs))
		conn.Close()
	}
}
