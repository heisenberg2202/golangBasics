package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple","Tomato","Peach","Pineapple"}
	fmt.Printf("Type of Fruitlist %T\n", fruitList)

	fruitList = append(fruitList, "Mango","Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)



	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 434
	highScores[2] = 334
	highScores[3] = 254

	highScores = append(highScores, 555,322,444)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	 fmt.Println(sort.IntsAreSorted(highScores))

	 //how to remove a value from slices based on index
	 //15/57

	 var courses = []string{"reactjs", "javascript","swift","python","ruby"}
	 
	 fmt.Println(courses)	
	 var index int = 2

	 courses = append(courses[:index] ,courses[index+1:]... )
	 fmt.Println(courses)

}