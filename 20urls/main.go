package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fhghfa434"

func main() {
	fmt.Println("Welcome to handling URL's in Golang")

	fmt.Println(myurl)

	//Parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)


	fmt.Println(qparams["coursename"])


	for _, val := range qparams{
		fmt.Println("Param is: ",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=ayush",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}