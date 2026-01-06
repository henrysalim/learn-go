package main

import "fmt"

func f(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("This is an integer")
	case string:
		fmt.Println("This is a string")
	default:
		fmt.Println("I don't know what this is")
	}
}

func main() {
	f("hello")
	f(42)
	f(false)
}
