package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list size is", len(fruitList))

	var vegList = [3]string{"potato","onion","beans"}
	fmt.Println("Veg list is ",vegList)
	
}