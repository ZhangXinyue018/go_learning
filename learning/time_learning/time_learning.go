package main

import (
	"time"
	"fmt"
)

func main()  {
	//currTime := time.Now()
	//fmt.Println(currTime)
	//fmt.Println(currTime.Hour())
	//fmt.Println(currTime.UTC())
	//fmt.Println(currTime.UTC().Hour())


	t := time.Now().UTC()
	nt := t.Format("2006-01-02T15:04:05.000")
	fmt.Println(nt)
}
