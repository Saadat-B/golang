package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://chaicode.com"

func main() {
	fmt.Println("LCO web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type %T\n", res)

	defer res.Body.Close()

	dataBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
