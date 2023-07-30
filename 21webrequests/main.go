package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
  "net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to webrequests in GO")

	// PerformGetRequest()

	//PerformPostJsonRequest()

  PerformPostFormRequest()

}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response, err  := http.Get(myurl)

	if err != nil{
		panic(err)
	}


	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder	
	content,_ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	
	fmt.Println("Bytecount is ",byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))


}


func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"


	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"ReactJs Course",
			"price":100,
			"platform":"udemy.com"
		}
	`)
	response, err := http.Post(myurl,"application/json", requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	content, _ :=ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest(){
  const myurl = "http://localhost:8000/postform"
  
  //formdata
  data := url.Values{}
  data.Add("firstname","ayush")
  data.Add("lastname","narayan")
  data.Add("email","ayush@go.dev")

  response, err := http.PostForm(myurl, data)
  
  if err != nil{
    panic(err)
  }
  defer response.Body.Close()

  content,_ := ioutil.ReadAll(response.Body)

  fmt.Println(string(content))

}
