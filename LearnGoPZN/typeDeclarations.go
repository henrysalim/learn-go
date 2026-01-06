package main

import "fmt"

func main()  {
	type NoKTP string

	var ktpJohn NoKTP = "12345678"
	fmt.Println(ktpJohn)
	fmt.Println(NoKTP("0987654321"))
}