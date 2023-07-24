package main

import "fmt"

func main() {

	defer fmt.Println("World") // will only run at the last
	defer fmt.Println("One")
	defer fmt.Println("Two") //LIFO
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i<5; i++{
		defer fmt.Print(i)
	}
}