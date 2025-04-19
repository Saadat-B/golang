package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"Tags,omitempty"`
}

func main() {

	fmt.Println("Welcome to the JSON video")
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"React", 299, "web", "abc123", []string{"new"}},
		{"Angular", 199, "web", "abc123", []string{"old"}},
		{"Nextjs", 399, "web", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
