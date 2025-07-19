package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostRequest()
}

// get data from route
func PerformGetRequest() {
	url := "http://localhost:5000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("message length is: ", response.ContentLength)
	fmt.Println("message status code is: ", response.StatusCode)

	content, err := ioutil.ReadAll(response.Body)
	fmt.Print(string(content))
	defer response.Body.Close()
}

func PerformPostRequest() {
	url := "http://localhost:5000/post"

	requestBody := strings.NewReader(`
		{
			"name":"Harshit Rajpurohit"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	defer response.Body.Close()

}
