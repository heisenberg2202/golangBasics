package main

import "fmt"

func main() {

	fmt.Println("Maps Examples")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "RUBY"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS is short form of",languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages",languages)

	//loops

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}


}