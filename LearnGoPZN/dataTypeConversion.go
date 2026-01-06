package main

import "fmt"

func main() {
	var nilai32 int32 = 32768 // maks = 32767, karna overflow maka balik ke nilai min
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var name = "Henry Salim"
	var e uint8 = name[1]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}