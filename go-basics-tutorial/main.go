package main

import (
	"errors"
	"fmt"
)

func main() {

	divideNum, err := divide(2, 0)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(divideNum)

}

func divide(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("divisor cannot be 0")
	}

	return a / b, nil
}
