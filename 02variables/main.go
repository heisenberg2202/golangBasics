package main

import "fmt"


const LoginToken string = "fdjdhsjfds"; // public variable -> first letter is capital



func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username )


	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn )


	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal )


	var val int = 255
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n",val )

	var smallFloat float32 = 255.4434333232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat )

	var largeFloat float64 = 255.4434333232
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n",largeFloat )


	//default values

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)


	//implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	

}