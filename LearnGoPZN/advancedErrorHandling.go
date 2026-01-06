package main

import "fmt"

type ValidationError struct {
	Field string
	Value interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s = %v", e.Field, e.Value)
}

func main() {
}
