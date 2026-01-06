package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Henry")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(message)
}
