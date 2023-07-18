package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of Golang")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))


	createdDate := time.Date(2020, time.February, 22, 23, 23,0,0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}