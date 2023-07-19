package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is ",result)

	proRes := proAdder(3,5,5,5,5)
	fmt.Println("Proresult is ",proRes)

	number, message := greeterTwo();
	fmt.Println("Values are",number,message)
 
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ...int) int{
	total := 0

	for _, val := range values{
		total += val
	}
	return total

}

func greeter(){
	fmt.Println("Namaste from Golang")
}

func greeterTwo() (int, string){
	return 4,"Hello"
}