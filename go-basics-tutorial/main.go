package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const geturl = "https://api.freeapi.app/api/v1/public/quotes/quote/random"
const posturl = "https://api.freeapi.app/api/v1/users/register"

func main() {
	fmt.Println("Welcome to web verb video ")

	// PerformGetRequest(geturl)
	// PerformPostRequest(posturl)
	PerformPostFormRequest(posturl)
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

func PerformPostRequest(myurl string) {

	requestBody := strings.NewReader(`

	{
 		"email": "saadatb.email@domain.com",
  		"password": "test@123",
  		"role": "ADMIN",
  		"username": "saadatb"
	}
	`)

	response, err := http.Post(posturl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, data, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Print(prettyJSON.String())
}

func PerformPostFormRequest(myurl string) {
	data := url.Values{}

	data.Add("email", "saadatbad.email@domain.com")
	data.Add("password", "ADMItestN")
	data.Add("username", "saadadtba")
	data.Add("role", "ADMIN")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, content, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Println(prettyJSON.String())

}
