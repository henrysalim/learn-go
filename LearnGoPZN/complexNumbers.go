package main

import (
	"fmt"
)

func main() {
	var a complex128 = 5 + 3i

	fmt.Println(a)
	fmt.Println(real(a))
	fmt.Println(imag(a))
}
