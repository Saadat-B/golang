package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func main() {
	fmt.Println("structs")


	firstUser := User{Name: "firstUser", Age: 5, IsActive: true}

	fmt.Printf("%#v",firstUser)

}
