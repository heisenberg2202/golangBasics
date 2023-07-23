package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang
	ayush := User{"Ayush","ayush@go.dev",true,22}
	fmt.Println(ayush)

	fmt.Printf("Ayush details are:  %+v\n", ayush)
	fmt.Printf("Name is %v and email is %v\n", ayush.Name,ayush.Email)

	ayush.GetStatus()
	ayush.newMail()
	fmt.Printf("Name is %v and email is %v\n", ayush.Name,ayush.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

func (u User) newMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:",u.Email)
}

