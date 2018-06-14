package main

import (
	"time"
	"fmt"
)

func main()  {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Println(currTime.Hour())
	fmt.Println(currTime.UTC())
	fmt.Println(currTime.UTC().Hour())
}
