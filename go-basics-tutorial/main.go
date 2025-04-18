package main

import (
	"fmt"
	"io"
	"net/http"
)

const myurl = "https://api.freeapi.app/api/v1/public/quotes/quote/random"

func main() {
	fmt.Println("Welcome to web verb video ")

	PerformGetRequest(myurl)
}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	content := string(data)

	fmt.Println(content)
	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length", response.ContentLength)

}
