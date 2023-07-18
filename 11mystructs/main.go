package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang
	ayush := User{"Ayush","ayush@go.dev",true,22}
	fmt.Println(ayush)

	fmt.Printf("Ayush details are:  %+v\n", ayush)
	fmt.Printf("Name is %v and email is %v", ayush.Name,ayush.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}