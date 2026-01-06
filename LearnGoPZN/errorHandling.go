package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0!")
	}

	return a / b, nil
}

func main() {
	res, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err.Error())
	}

	println("Result:", res)
}
