package main

import "fmt"

func main() {

	var array [5]string

	array[0] = "firstUser"
	array[1] = "secondUser"
	array[2] = "thirdUser"
	array[3] = "fourthUser"
	array[4] = "fifthUser"

	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	myMap := make(map[string]int)

	myMap["first"] = 1

	fmt.Println(myMap)
}
