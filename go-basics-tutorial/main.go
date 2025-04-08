package main

import "fmt"

func main() {

	message := greet("Saadat")
	fmt.Println(message)
}

func greet(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}
