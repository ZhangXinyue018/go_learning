package main

import (
	"fmt"
	"net/http"
)


func main() {
	go http.HandleFunc("/user", userHandleFunc)
	go http.HandleFunc("/test", testHandleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}

func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Calling /user using method: %v \n", r.Method)
	w.Write([]byte("hello"))
}

func testHandleFunc(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Calling /test using method: %v \n", r.Method)
	w.Write([]byte("hello test"))
}
