package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file. Let's go!!"

	file, err := os.Create("./mygofile.txt")

	//  

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mygofile.txt")
}



func readFile(filename string){

	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}

}