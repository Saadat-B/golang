package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}
type Robot struct {
	Name     string
	Age      int
	IsActive bool
}

type Greet interface {
	Greeting() string
}

func (u User) Greeting() string {
	return "Hello I'm " + u.Name + " aged about " + strconv.Itoa(u.Age) + "years\n"
}
func (r Robot) Greeting() string {
	return "Hello I'm " + r.Name + " aged about " + strconv.Itoa(r.Age) + "years\n"
}

func printGreet(g Greet) string {
	greeting := g.Greeting()
	return greeting
}

func main() {
	fmt.Println("structs")

	firstUser := User{Name: "firstUser", Age: 5, IsActive: true}
	firstRobot := Robot{Name: "firstRobot", Age: 10, IsActive: true}

	// fmt.Printf("%#v", firstUser)
	// fmt.Printf("%#v", firstRobot)

	fmt.Print(printGreet(firstUser))
	fmt.Print(printGreet(firstRobot))

}
