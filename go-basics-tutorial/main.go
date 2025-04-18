package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hsdif4343"

func main() {

	// fmt.Println("Welcome to handling URLS in golang")
	fmt.Println(myurl)

	//parsing

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Query())
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Println(key, val)
	}

}
