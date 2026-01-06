package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var stringToInt string = "12345"
	intValue, err := strconv.Atoi(stringToInt)

	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Converted integer value:", intValue)
	}
	fmt.Println()

	var float64Value float64 = 9.80
	var negValue = int(-42)
	fmt.Println("Floor for float64Value:", math.Floor(float64Value))
	fmt.Println("Ceiling for float64Value:", math.Ceil(float64Value))
	fmt.Println("Nearest value for float64Value:", math.Round(float64Value))
	fmt.Println("Absolute value for negValue:", math.Abs(float64(negValue)))

	fmt.Println()
	var fullName string = "Henry Salim"
	fmt.Println("Fullname contains 'lim':", strings.Contains(fullName, "lim")) //case sensitive
	fmt.Println("Index of e:", strings.Index(fullName, "e")) // -1 if not found
	fmt.Println("Split the string", strings.SplitAfter(fullName, "Sa"))
	fmt.Println("Remove 'H' and 'lim':", strings.Trim(fullName, "Hlim")) // remove leading and trailing characters
}
