package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ",ptr)

	myNumber := 23
	// Refrence -> &
	var ptr = &myNumber 
	fmt.Println("Value of actual pointer is ",ptr)
	fmt.Println("Value of actual pointer is ",*ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ",*ptr)

}