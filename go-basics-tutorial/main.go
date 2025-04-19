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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
  "name": "Angular",
  "price": 199,
  "Platform": "web",
  "Tags": [ "old" ]
 	}
 `)

	var lcoCourse course

	isValidJson := json.Valid(jsonDataFromWeb)

	if isValidJson {
		fmt.Println("JSON is valid")
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// var myOnlineData map[string]interface{} can also use 'any' like below which is an alias for interface{}
	var myOnlineData map[string]any

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("myonlinedata %#v", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("the Key is %v and the value is %v and the type is %T\n", k, v, v)
	}

}
